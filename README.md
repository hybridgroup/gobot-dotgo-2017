# dotGO 2017 workshop

This is the code and activities for the dotGo and Golang Paris hardware programming workshop using Gobot ([https://gobot.io](https://gobot.io))

There are three workshop activities:

    - Sensor station
    - Rover
    - Drone

Each activity consists of a series of steps to end up with the complete solution, with a device controlled by Go using Gobot, that communicates data back to an [Intel IoT Gateway](https://software.intel.com/en-us/iot/hardware/gateways) that is running the base station software.

You can do the activities in any order, for example first "rover" then "Sensor station", then "Drone". Each step within the activity should be done in sequence.

## Sensor station

This activity uses an Arduino along with the Intel Grove Starter Kit to create a sensor station for our interplanetary explorations. The sensor station communicates data back to the base station.

## Rover

This activity uses a Sphero SPRK+ or Sphero Ollie as part of a remote rover system to explore the planet's surface, and to return data back to the base station.

## Drone

This activity uses a Parrot Minidrone to provide air support for our interplanetary expedition. The drone communicates back to the base station.

## Base station

Uses an Intel IoT gateway running a Gobot program and the Mosquito MQTT machine to machine messaging server ([https://mosquitto.org/](https://mosquitto.org/)). This provides the central "base station" used by all of the workshop activities.
