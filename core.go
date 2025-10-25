package main

import (
	"fmt"
	"os"
	"strconv"

	libLWD "github.com/Hydriam/LinuxWorkshopDownloader/LibLWD"
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.Hydriam.LinuxWorkshopDownloader", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	// get the ui file
	builder := gtk.NewBuilderFromFile("lwd-ui.ui")
	// When we get mainWindow are childs work but they dont have variables assigned to them
	// So for some widgets we need to get them separately
	mainWindow := builder.GetObject("mainWindow").Cast().(*gtk.Window)
	mainWindow.SetDefaultSize(400, 200)
	// entry with mod app id
	modAppID := builder.GetObject("modAppID").Cast().(*gtk.Entry)
	// entry with game app id
	gameAppID := builder.GetObject("gameAppID").Cast().(*gtk.Entry)
	// list of the mods to download
	modList := builder.GetObject("modList").Cast().(*gtk.ListView)
	// the string list for the mod list
	modListStrings := builder.GetObject("modListStrings").Cast().(*gtk.StringList)
	// mod list factory
	// we cant use listItemFactory from .ui, we need to create and set it up in code
	/* note to myself : https://github.com/Tom5521/gtk4tools
	   has an example of this listview spaghetti*/
	modListFactory := gtk.NewSignalListItemFactory()
	modListFactory.ConnectSetup(func(listitem *glib.Object) {
		listitem.Cast().(*gtk.ListItem).SetChild(gtk.NewLabel(""))
	})
	modListFactory.ConnectBind(func(listitem *glib.Object) {
		obj := listitem.Cast().(*gtk.ListItem).Item().Cast().(*gtk.StringObject)
		listitem.Cast().(*gtk.ListItem).Child().(*gtk.Label).SetText(obj.String())
	})
	// add button
	addButton := builder.GetObject("addButton").Cast().(*gtk.Button)
	addButton.ConnectClicked(func() {
		/* to avoid confusion:
		modAppID - The gtk entry where you type appID of the mod
		appID - The text we retrive from modAppID
		modListString - The list */
		appID := modAppID.Text()
		// TODO: instead of returning display error dialogs
		if appID == "" {
			return
		}
		// check if the appid is a number
		_, err := strconv.Atoi(appID)
		if err != nil {
			fmt.Println("error: appID needs to be a number")
			return
		}

		modListStrings.Append(appID)
		modAppID.SetText("")

	})
	// download button
	downloadButton := builder.GetObject("downloadButton").Cast().(*gtk.Button)
	downloadButton.ConnectClicked(func() {
		exists := libLWD.CheckSteamcmd()
		if exists == false {
			alertDialog := adw.NewAlertDialog("Steamcmd not installed", "Would you like the program to install it?")
			alertDialog.AddResponse("yes", "Yes")
			alertDialog.AddResponse("no", "No")
			alertDialog.SetDefaultResponse("yes")
			alertDialog.SetCloseResponse("no")
			alertDialog.ConnectResponse(func(response string) {
				if response == "yes" {
					go libLWD.GetSteamcmd()
				} else {

				}
			})
			alertDialog.Present(mainWindow)
			return
			// CRITICAL TODO: make the program wait till getSteamcmd ends and then downloads mods, now when just get steacmd without downloading mods
		}
		if gameAppID.Text() == "" {
			fmt.Println("Error: gameAppID is empty")
			return
		}
		var workshopIDs []string
		for i := 0; i < int(modListStrings.NItems()); i++ {
			workshopIDs = append(workshopIDs, modListStrings.String(uint(i)))
		}
		//fmt.Println(workshopIDs)
		go libLWD.DownloadFromSteamcmd(gameAppID.Text(),
			workshopIDs,
			true)
	})
	// set modListFactory as the factory for mod list
	modList.SetFactory(&modListFactory.ListItemFactory)
	// display the app
	mainWindow.SetApplication(app)
	mainWindow.Show()
}
