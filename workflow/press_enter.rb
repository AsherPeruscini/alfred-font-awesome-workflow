#!/usr/bin/env ruby
# encoding: utf-8

require './require_helper'

argv = FontAwesome.argv(ARGV)
print FontAwesome.css_class_name(argv.icon_id)
