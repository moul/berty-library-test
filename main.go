package main

import (
	"context"
	"flag"
	"log"
	"os"

	"berty.tech/berty/v2/go/pkg/bertymessenger"
	"github.com/gogo/protobuf/proto"
	qrterminal "github.com/mdp/qrterminal/v3"
	"google.golang.org/grpc"
	"moul.io/u"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}

var nodeAddr = flag.String("addr", "127.0.0.1:9091", "remote 'berty daemon' address")

func run(args []string) error {
	flag.Parse()
	if len(args) > 0 {
		return flag.ErrHelp
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// open gRPC connection to the remote 'berty daemon' instance
	var client bertymessenger.MessengerServiceClient
	{
		conn, err := grpc.Dial(*nodeAddr, grpc.WithInsecure())
		checkErr(err)
		client = bertymessenger.NewMessengerServiceClient(conn)
	}

	// get sharing link
	{
		req := &bertymessenger.InstanceShareableBertyID_Request{DisplayName: "berty-library-test"}
		res, err := client.InstanceShareableBertyID(ctx, req)
		checkErr(err)
		log.Printf("berty id: %s", res.HTMLURL)
		qrterminal.GenerateHalfBlock(res.HTMLURL, qrterminal.L, os.Stdout)
	}

	// event loop
	{
		s, err := client.EventStream(ctx, &bertymessenger.EventStream_Request{})
		checkErr(err)

		go func() {
			for {
				gme, err := s.Recv()
				checkErr(err)

				// parse event's payload
				update, err := gme.Event.UnmarshalPayload()
				checkErr(err)

				switch gme.Event.Type {
				case bertymessenger.StreamEvent_TypeContactUpdated:
					// auto-accept contact requests
					contact := update.(*bertymessenger.StreamEvent_ContactUpdated).Contact
					log.Printf("<<< %s: contact=%q conversation=%q name=%q", gme.Event.Type, contact.PublicKey, contact.ConversationPublicKey, contact.DisplayName)
					if contact.State == bertymessenger.Contact_IncomingRequest {
						req := &bertymessenger.ContactAccept_Request{PublicKey: contact.PublicKey}
						_, err = client.ContactAccept(ctx, req)
						checkErr(err)
					}

				case bertymessenger.StreamEvent_TypeInteractionUpdated:
					// auto-reply to users' messages
					interaction := update.(*bertymessenger.StreamEvent_InteractionUpdated).Interaction
					log.Printf("<<< %s: conversation=%q", gme.Event.Type, interaction.ConversationPublicKey)
					if interaction.Type == bertymessenger.AppMessage_TypeUserMessage && !interaction.IsMe && !interaction.Acknowledged {
						userMessage, err := proto.Marshal(&bertymessenger.AppMessage_UserMessage{
							Body: "Hey.",
						})
						checkErr(err)

						_, err = client.Interact(ctx, &bertymessenger.Interact_Request{
							Type:                  bertymessenger.AppMessage_TypeUserMessage,
							Payload:               userMessage,
							ConversationPublicKey: interaction.ConversationPublicKey,
						})
						checkErr(err)
					}

				default:
					log.Printf("<<< %s: ignored", gme.Event.Type)
				}
			}
		}()
	}

	u.WaitForCtrlC()
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
