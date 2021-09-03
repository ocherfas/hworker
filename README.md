# hworker - Intro
A tool to help you stay healthy while working long hours on the computer.

## How did I come up with this?
After a visit to my optometrist and after asking him for suggestions on how to cope with my decreased vision, he suggested to have a break every few minutes to let my eyes rest.
I find it hard to remeber to do such a thing when working with full concentraction on something. So like a good developer - I decided to do something about it.

# Current features
- Capture every event from mouse or keyboard to capture "Computer activity".
- Present a desktop notification after a long period of continious activity.
- Configuration for how much is it too long.
- Configuration for inactivity timeout.

# Getting Started
**Please note that this application was only built and tested on Linux**

Having said that - Should not encounter that many problems with mac or windows, so any contribution will be welcomed.

## Build
`./hack/build.sh`
## Install
`sudo ./hack/install.sh`
## Configuration
Installation will install the application on `/usr/local/hworker/`. There you can find the file `config.yaml` and in it you will find the configuration options:

```
activity-monitors:
  - max-activity-time: 15s
    inactivity-time: 5s
    message-format: message1 %s
  - max-activity-time: 30s
    inactivity-time: 5s
    message-format: message2 %s
```

`max-activity-time:` # After how much time should pop up the notification.

`inactivity-time:` # Inactivity timeout. If user did not have any  activity during that time, the timer for the notification will be reset.

`message-format:` # The message that will be shown to the user for continious activity. Use %s to place the `max-activity-time` value inside the message.

# Contact me
Feel free to contact me for any suggestion, contribution, question, or anything else.

Contact me on: ocherfas@gmail.com