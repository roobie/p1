#!/usr/bin/env luajit

local args = ... or ""

print(args)

local function help ()
  print([[ctl.lua
  Usage:
    ctl.lua build|run
]])
end

local function build ()
end

if #args < 1 then
  help()
  os.exit(1)
end
