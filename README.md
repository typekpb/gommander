# gommander

idea:

event types:
* https://github.com/google/go-github/blob/master/github/event_types.go

uis are:
* event handlers/listeners/observers
* cmd triggers

cmds are:
* event publisher

event/event_dispatcher.go responsible for:
* event triggering
* listener registration

tasks are:
* created once event_* triggered
* live until task runs or is cancelled
* trigger progress notification events

workflow:
* cmd -> triggers event
* task is an observer for specific event (or cmd?)
* dispatcher calls all the registered listeners

for process dispatch + output reading, used:
* https://github.com/go-cmd/cmd

rules for the tasks:
* golang libs
* commands (e.g. unzip/)

releases:
https://github.com/wangyoucao577/go-release-action