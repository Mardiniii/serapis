require 'chronic'

time = Time.now
puts Time.now
puts Chronic.parse('tomorrow')

puts "Insert value for A:"
a = gets.chomp.to_i
puts "A value is: #{a}"

puts "Insert value for B:"
b = gets.chomp.to_i
puts "B value is: #{b}"

sum = a + b
puts "A + B result is: #{sum}"
