package command

import (
	"fmt"
	"github.com/itzngga/goRoxy/internal/handler"
	"github.com/itzngga/goRoxy/util"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func UrlImageButtonCommand() {
	AddCommand(&handler.Command{
		Name:        "url-button-image",
		Description: "Url Image button command",
		Category:    handler.UtilitiesCategory,
		RunFunc:     UrlImageButtonRunFunc,
	})
}

func UrlImageButtonRunFunc(c *whatsmeow.Client, args handler.RunFuncArgs) *waProto.Message {
	id, _ := args.Locals.Load("uid")

	image, err := util.UploadImageFromUrl(c, "https://picsum.photos/id/870/1280/720?grayscale&blur=2", "Testing")
	if err != nil {
		fmt.Printf("\nError uploading image :  %v", err)
		return nil
		// remember to return nil when error is returned
	}

	return util.CreateImageButton("test", "footer",
		&waProto.ButtonsMessage_ImageMessage{
			ImageMessage: image,
		},
		util.GenerateButton(id, "!help", "Help"),
	)
}
