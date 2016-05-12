require 'formula'

class CwlogsCat < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/cwlogs-cat'
  url "https://github.com/winebarrel/cwlogs-cat/releases/download/v#{VERSION}/cwlogs-cat-v#{VERSION}-darwin-amd64.gz"
  sha256 'f22dfd3e671ade54a8e7a1e7d41c3cf3adfb6654ae6554b54e30924fa6e5ebe3'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-cat.git', :branch => 'master'

  def install
    system "mv cwlogs-cat-v#{VERSION}-darwin-amd64 cwlogs-cat"
    bin.install 'cwlogs-cat'
  end
end
