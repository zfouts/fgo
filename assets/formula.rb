#
# assets/formula.rb
#
# Copyright (c) 2016 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
class {{.Package | Title}} < Formula
  desc ""
  homepage "https://github.com/{{.UserName}}/{{.Package}}"
  version "{{"{{.Version}}"}}"

  if Hardware::CPU.is_64_bit?
    url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.FileName64}}"}}"
    sha256 "{{"{{.Hash64}}"}}"
  else
    url "https://github.com/{{.UserName}}/{{.Package}}/releases/download/v{{"{{.Version}}"}}/{{"{{.FileName386}}"}}"
    sha256 "{{"{{.Hash386}}"}}"
  end

  def install
    bin.install "{{.Package}}"
  end

  test do
    system "{{.Package}}"
  end

end
