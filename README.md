# haSystray

This application is a system tray icon for Home Assistant. It allows you to control your Home Assistant instance from the system tray.

> [!IMPORTANT]  
> This project is not affiliated with Home Assistant or any of its developers in any way.
> This project only uses Webhooks to trigger automations in Home Assistant and could be used for any other webhooks you'd like.

> [!CAUTION]
> This application is just a proof of concept. Use at your own risk!
> Do not use this application to control any automations that could cause damage to you, any other creatures (including humans and other species).
> Do not use this application to control any automations that could cause damage to your property or the property of others.

## Thanks!

Thanks to [@getlantern](https://github.com/getlantern) on GitHub to make this project possible. I used his [systray](https://github.com/getlantern/systray) library to create this application.

## Configuration

The configuration is done via a `config.json` file in the same directory as the application.

The following is an example configuration file for this:

![video of the app running](assets/example_screenshot.gif)

```json
{
  "programTooltip": "HaSystray",
  "actions": [
    {
      "name": "Light",
      "tooltip": "Turn on the light",
      "webhook": "xxx",
      "data": {
        "scene": "nacht"
      },
      "type": "POST",
      "actions": [
        {
          "name": "Nacht",
          "tooltip": "Turn off the light",
          "webhook": "xxx",
          "data": {
            "scene": "nacht"
          },
          "type": "POST"
        },
        {
          "name": "Hell",
          "tooltip": "Turn off the light",
          "webhook": "xxx",
          "data": {
            "scene": "hell"
          },
          "type": "POST"
        },
        {
          "name": "Schreibtisch",
          "tooltip": "Turn off the light",
          "webhook": "xxx",
          "data": {
            "scene": "schreibtisch"
          },
          "type": "POST"
        }
      ]
    }
  ]
}
```

The corresponding Automation for Home Assistant would look something like this:
```yaml
alias: HaSystray Exmaple
description: ""
trigger:
  - platform: webhook
    allowed_methods:
      - POST
    local_only: true
    webhook_id: "xxx"
condition: []
action:
  - service: scene.turn_on
    target:
      entity_id: >-
        {% if trigger.json.scene == "nacht" %} scene.nacht {% elif
        trigger.json.scene == "hell" %} scene.hell {% elif
        trigger.json.scene == "schreibtisch" %} scene.schreibtisch {% endif %}
    metadata: {}
mode: single
```

## Features

This application is very simple and is only supposed to trigger Automations in Home Assistant via the Webhook API. 
It does not support any other features of Home Assistant for the time being.

> [!WARNING]
> The Homeassistant developers say, that the Webhook API is not secure (look [here](https://www.home-assistant.io/docs/automation/trigger/#webhook-security) for reference on this). 
> So use this feature at your own risk!

## Supported Platforms

This application is written and tested to work on macOS. Contributions to support other platforms are welcome.

## Contributing

This is only supposed to be used for my personal use. But if you want to contribute, feel free to do so.

## [License](LICENSE.md)

This project is licensed under the Apache 2 License - see the [LICENSE.md](LICENSE.md) file for details

