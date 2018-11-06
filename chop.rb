#!/usr/bin/env ruby
# encoding: UTF-8

def usage(message)
  name = File.basename($PROGRAM_NAME)

  if message
    puts "#{name}: #{message}"
    puts
  end

  puts <<-eos
#{name} --header <HEADER> --wanted <WANTED> <FILE1> <FILE2> ... <FILEN>
    Processes all the files <FILE1> to <FILEN> and breaks them into blocks on
    any line containing the <HEADER> text. If the block also contains the
    <WANTED> text the block will be displayed to stdout

    Both --header and --wanted are required arguments
eos

  exit(1)
end

def find_opts(list, required = [])
  rest = []
  args = {}

  while list.any?
    x = list.shift
    if x.index('--') == 0
      key = x[2..-1].downcase
      usage("Argument #{x} already supplied") if args.key?(key)

      if key.include?('=')
        key, value = key.split('=', 2)
      else
        value = list.shift
      end

      usage("No value given for #{x}") if value.nil?
      args[key] = value
    else
      rest << x
    end
  end

  required.each do |r|
    usage("Required argument --#{r} is missing") unless args.key?(r)
  end

  return args, rest
end

args, rest = find_opts(ARGV, %w(wanted header))

def process(text, wanted)
  return unless text.include?(wanted)
  puts text
  puts
end

text = ''

rest.each do |filename|
  File.open(filename, 'r').each do |line|
    if line.include?(args['header'])
      process(text, args['wanted'])
      text = ''
    end
    text << line
  end
  process(text, args['wanted'])
end
