#!/bin/sh

test -f BemfaZhilian-Installer.dmg && rm BemfaZhilian-Installer.dmg

wails build -u -platform darwin/universal
create-dmg \
    --volname "BemfaZhilian Installer" \
    --volicon "build/bin/BemfaZhilian.app/Contents/Resources/iconfile.icns" \
    --icon-size 75 \
    --window-size 600 400 \
    --icon "BemfaZhilian.app" 200 170 \
    --app-drop-link 400 170 \
    --hide-extension "BemfaZhilian.app" \
    "BemfaZhilian-Installer.dmg" \
    "build/bin/"
