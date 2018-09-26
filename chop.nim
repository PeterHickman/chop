import strutils
import os
import tables

var args = initTable[string, string]()
var files: seq[string] = @[]
var text: seq[string] = @[]

proc get_args() =
  var i = 1
  var line, value: string

  while i <= paramCount():
    line = paramStr(i).string

    if line.startsWith("--"):
      line = toLowerAscii(line[2 .. ^1])
      i += 1
      value = paramStr(i).string
      args[line] = value
    else:
      files.add(line)

    i += 1

proc process(text: seq[string]) =
  var found = false

  for line in text:
    if contains(line, args["wanted"]):
      found = true
      break

  if found:
    for line in text:
      echo line
    echo ""

proc add_line(line: string) =
  if contains(line, args["header"]):
    process(text)
    text = @[]
  text.add(line)

get_args()

if args.len() != 2:
  echo "There are 2 required argument --header and --wanted"
  quit(QuitFailure)

if not hasKey(args, "header"):
  echo "Missing argument --header"
  quit(QuitFailure)

if not hasKey(args, "wanted"):
  echo "Missing argument --wanted"
  quit(QuitFailure)

if files.len() == 0:
  for line in stdin.lines:
    add_line(line)
  process(text)
else:
  for file in files:
    for line in lines file:
      add_line(line)
    process(text)

