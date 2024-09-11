a simple text to speech program that outputs spoken text patched together from a set of audio files.

give it an input of a string to turn to speech, and a dataset of timestamped lyrics and audiofiles.
and receive a mosaic of audio snipits to speak your input string.



# todo
- process accurate timestamped lyrics from audio files with [whisper-timestamped](https://github.com/linto-ai/whisper-timestamped)
- get a list of start-end times of audio files to patch together
- splice the specific parts of the audio files we want together
- ~~profit~~
