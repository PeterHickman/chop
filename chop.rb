#!/usr/bin/env ruby
# encoding: UTF-8

def find_opts(list, required = [])
  rest = []
  args = {}

  while list.any?
    x = list.shift
    if x.index('--') == 0
      key = x[2..-1].downcase
      raise "Argument #{x} already supplied" if args.key?(key)

      if key.include?('=')
        key, value = key.split('=', 2)
      else
        value = list.shift
      end

      raise "No value given for #{x}" if value.nil?
      args[key] = value
    else
      rest << x
    end
  end

  required.each do |r|
    raise "Required argument --#{r} is missing" unless args.key?(r)
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
