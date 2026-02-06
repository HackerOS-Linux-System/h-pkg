require "file_utils"

def main
  home = ENV["HOME"]
  config_dir = "#{home}/.config/h-pkg"
  config_file = "#{config_dir}/config.hacker"
  default_mode = "cli"
  mode = default_mode

  if File.exists?(config_file)
    content = File.read(config_file).strip
    if content.starts_with?("[") && content.ends_with?("]")
      mode = content[1...-1].downcase
    end
  end

  bin_path = "#{home}/.hackeros/h-pkg/h-pkg-#{mode}"

  unless File.exists?(bin_path)
    puts "Binary not found: #{bin_path}"
    exit(1)
  end

  args = ARGV
  Process.exec(bin_path, args)
end

main
