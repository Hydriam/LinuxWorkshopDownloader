# LinuxWorkshopDownloader
A Linux Gui App For Downloading Steam Workshop Elements

**The app is still work-in-progress**
## What this does?
This tool can download mods from steam workshop using steamcmd,
This is uselful if you have a game on gog which most modding community is on steam workshop
**You need 32 bit version of glibc installed beacuse steamcmd needs it**
## Usage
In the mod appid entry add appids of mods you want to download, you can find mod appid by looking at the link in your browser, then fill the game app id entry with the app id but of the game the mod belongs to.
First time downloading the app will display a "Steamcmd not installed" dialog, click yes and the app will download steamcmd, then just click download button again and it should download the mods.
Downloaded mods should be under ~/.local/share/Steam/steamapps/workshop/, if not check steamcmd.log
## Building
This app is writen in golang, uses gotk4 and cambalache for gui, and codeclysm/extract for extracting steamcmd.
To build run:
```
git clone https://github.com/Hydriam/LinuxWorkshopDownloader/
cd LinuxWorkshopDownloader
go build -o LinuxWorkshopDownloader
```
You need golang installed on your linux system.
