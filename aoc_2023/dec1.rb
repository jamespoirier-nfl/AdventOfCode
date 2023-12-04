#!/usr/bin/env ruby

# Advent of Code - Day 1

class AdventOfCode
  def initialize
    @input_lines = []
    # Initialize any other variables or data structures here
    puts "🐹🐹🐹🐹 Initializing..."
    @running_total_pt_1 = 0
    @running_total_pt_2 = 0
  end

  def read_input(file_path)
    # Read input lines from the specified file
    File.readlines(file_path, chomp: true).each do |line|
      @input_lines << line
    end
  end

  def process_input
    # Process each input line for Part One
    @input_lines.each do |line|
      solve_part_one(line)
    end

    # Process each input line for Part Two
    @input_lines.each do |line|
      solve_part_two(line)
    end
  end


  def solve_part_one(line)
    # Extract the first and last single-digit numbers from each line
    first_number = line.match(/\d/).to_s.to_i
    last_number = line.scan(/\d/).last.to_i

    # Combine the first and last numbers by concatenating them
    combined_number = "#{first_number}#{last_number}".to_i

    # puts "Processed line: #{line}"
    # puts "First number: #{first_number}, Last number: #{last_number}, Combined number: #{combined_number}"
    # Add the combined number to the running total
    @running_total_pt_1 += combined_number
    # puts "Running total: #{@running_total_pt_1}"

  end

  def solve_part_two(line)
    # Initialize a hash to map words to their corresponding digits
    word_to_digit = {
      'one' => 1,
      'two' => 2,
      'three' => 3,
      'four' => 4,
      'five' => 5,
      'six' => 6,
      'seven' => 7,
      'eight' => 8,
      'nine' => 9
    }

    first_word_or_digit = line.match(/(?:one|two|three|four|five|six|seven|eight|nine|\d)/).to_s

    first_digit = word_to_digit[first_word_or_digit] || first_word_or_digit.to_i

    last_word_or_digit = line.reverse.scan(/(?:eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d)/).first.to_s.reverse

    last_digit = word_to_digit[last_word_or_digit] || last_word_or_digit.to_i
    # puts "Processed line: #{line}, Last word or digit: #{last_word_or_digit}, Last digit: #{last_digit}"

    # Combine the first and last digits by concatenating them
    combined_number = "#{first_digit}#{last_digit}".to_i
    # puts "Combined number: #{combined_number}"

    # Add the combined number to the running total
    @running_total_pt_2 += combined_number
    # puts "Running total: #{@running_total_pt_2}"

  end
end

if __FILE__ == $PROGRAM_NAME
  # Check if the input file path is provided as a command-line argument
  if ARGV.empty?
    puts "Usage: #{$PROGRAM_NAME} <input_file_path>"
    exit(1)
  end

  file_path = ARGV[0]
  advent_of_code = AdventOfCode.new
  advent_of_code.read_input(file_path)
  advent_of_code.process_input

  # Print the running total for Part One
  puts "Running total for Part One: #{advent_of_code.instance_variable_get(:@running_total_pt_1)}"

  # Print the running total for Part Two
  puts "Running total for Part Two: #{advent_of_code.instance_variable_get(:@running_total_pt_2)}"

end
