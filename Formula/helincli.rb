class Helincli< Formula
    desc "何林测试brew cli的工具"
    url "https://github.com/username/helincli/releases/download/v0.1.0/helincli_0.1.0_Darwin_x86_64.tar.gz"
    version "alpha1"
    
    if OS.mac?
        @os="macos"
        if Hardware::CPU.intel?
            @arch = "amd64"
            sha256 "todo"
        else
            @arch = "arm64"
            sha256 "todo"
        end
    elsif OS.linux?
        @os = "linux"
        if Hardware::CPU.intel?
            @arch = "amd64"
            sha256 "todo"
        end
    end
    def install
        bin.install "helincli"
        FileUtils.chmod 0755, bin/"helincli"
    end
