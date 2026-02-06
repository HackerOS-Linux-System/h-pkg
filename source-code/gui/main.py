import gi
gi.require_version('Gtk', '4.0')
from gi.repository import Gtk, Gio
import subprocess
import sys

def run_apt(command, args):
    full_cmd = ['sudo', 'apt'] + command + args
    subprocess.run(full_cmd)

def on_install(button):
    # For simplicity, assume package name from entry or hardcoded; extend as needed
    package = entry.get_text()
    run_apt(['install', '-y'], [package])

def on_remove(button):
    package = entry.get_text()
    run_apt(['remove', '-y'], [package])

def on_update(button):
    run_apt(['update'], [])

def on_upgrade(button):
    run_apt(['upgrade', '-y'], [])

def on_search(button):
    term = entry.get_text()
    result = subprocess.run(['apt', 'search', term], capture_output=True, text=True)
    textbuffer = textview.get_buffer()
    textbuffer.set_text(result.stdout)

def main(argv):
    global entry, textview

    app = Gtk.Application(application_id='org.example.hpkg')
    app.connect('activate', on_activate)
    app.run(sys.argv)

def on_activate(app):
    window = Gtk.ApplicationWindow(application=app)
    window.set_title('H-PKG GUI')
    window.set_default_size(400, 300)

    box = Gtk.Box(orientation=Gtk.Orientation.VERTICAL, spacing=6)
    window.set_child(box)

    entry = Gtk.Entry()
    box.append(entry)

    install_btn = Gtk.Button(label='Install')
    install_btn.connect('clicked', on_install)
    box.append(install_btn)

    remove_btn = Gtk.Button(label='Remove')
    remove_btn.connect('clicked', on_remove)
    box.append(remove_btn)

    update_btn = Gtk.Button(label='Update')
    update_btn.connect('clicked', on_update)
    box.append(update_btn)

    upgrade_btn = Gtk.Button(label='Upgrade')
    upgrade_btn.connect('clicked', on_upgrade)
    box.append(upgrade_btn)

    search_btn = Gtk.Button(label='Search')
    search_btn.connect('clicked', on_search)
    box.append(search_btn)

    textview = Gtk.TextView()
    box.append(textview)

    window.present()

if __name__ == '__main__':
    main(sys.argv)
