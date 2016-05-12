require 'formula'

class CwlogsCat < Formula
  VERSION = '0.1.1'

  homepage 'https://github.com/winebarrel/cwlogs-cat'
  url "https://github.com/winebarrel/cwlogs-cat/releases/download/v#{VERSION}/cwlogs-cat-v#{VERSION}-darwin-amd64.gz"
  sha256 'e81020b0c8059359d9b36dc9b967c41bca2f0b1a62139fe0d4979a833a2179af'
  version VERSION
  head 'https://github.com/winebarrel/cwlogs-cat.git', :branch => 'master'

  def install
    system "mv cwlogs-cat-v#{VERSION}-darwin-amd64 cwlogs-cat"
    bin.install 'cwlogs-cat'
  end
end
