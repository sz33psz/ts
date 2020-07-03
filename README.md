TS - simple timestamp tool:

`ts` - print current timestamp
`ts -1d12h` - print unix timestamp from 1 day and 12 hours ago
`ts 0h0m0s` - print unix timestamp of beginning of current day
`ts +1d 0h0m0s` - print timestamp of tomorrow's midnight
`ts 12h30m` - print unix timestamp of 12:30:00 today
`ts -t` - print current ISO 8601 UTC time
`ts -t 1590000000` - print ISO 8601 UTC time from unix timestamp
`ts -t 1590000000 +1d` - print ISO 8601 UTC time from 1 day after the unix timestamp