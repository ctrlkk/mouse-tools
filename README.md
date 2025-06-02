
# 构建步骤

系统: Ubuntu24

### GoHook
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev

### C
sudo apt-get install mingw-w64

### 构建 Windows 64位
GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 wails build -platform windows/amd64
