
install:
	 @echo "==> fetchfetch installer"
	 @echo "==> Made by gentoo-btw"
	 @echo " "
	 @echo "==> Installing fetchfetch.."
	 @test -f /usr/bin/go && echo "==> Found go executable at /usr/bin/go."
	 @test -f /usr/bin/go || echo "==> Go not found. please install go." || exit 1
	 @echo "==> Running go to install fetchfetch.."
	 @go install github.com/NikonP/fetchfetch@latest
	 @echo "==> Installed fetchfetch."
	 @echo "==> Make sure ~/go/bin is in your PATH"
	 @echo " "
	 @echo "==> Also. stop trying to make fetch happen (if your gonna make it happen)"
	 @echo "==> We are tired of new fetch tools over and over again. so please don't you dare publish your fetch tool into reddit (e.g r/linux)"

update:
	 @echo "==> fetchfetch installer"
	 @echo "==> Made by gentoo-btw"
	 @echo " "
	 @echo "==> Updating fetchfetch"
	 @echo "==> Removing current fetchfetch executable. assuming path ~/go/bin/"
	 @rm ~/go/bin/fetchfetch
	 @echo "==> Installing fetchfetch (again)"
	 @go install github.com/NikonP/fetchfetch@latest
	 @echo "==> Updated fetchfetch"
	 @echo "==> Make sure ~/go/bin is in your PATH."
	 @echo " "
	 @echo "==> Also. stop trying to make fetch happen (if your gonna make it happen)"
	 @echo "==> We are tired of new fetch tools over and over again. so please don't you dare publish your fetch tool into reddit (e.g r/linux)"

uninstall:
	 @echo "==> fetchfetch installer"
	 @echo "==> Made by gentoo-btw"
	 @echo " "
	 @echo "==> Removing fetchfetch. assuming to remove ~/go/bin/fetchfetch"
	 @rm ~/go/bin/fetchfetch
	 @echo "==> Thanks for using fetchfetch"
	 @echo " "
	 @echo "==> Also. stop trying to make fetch happen (if your gonna make it happen)"
	 @echo "==> We are tired of new fetch tools over and over again. so please don't you dare publish your fetch tool into reddit (e.g r/linux)"
