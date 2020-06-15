TS - simple timestamp tool:

`ts` - print current timestamp
`ts -m` - print current timestamp in milliseconds (applies to all next examples)
`ts -1d12h` - print unix timestamp from 1 day and 12 hours ago
`ts 0h0m0s` - print unix timestamp of beginning of current day
`ts +1d 0h0m0s` - print timestamp of beginning of tomorrow
`ts 12h30m` - print unix timestamp of 12:30:00 today
`ts 1577836800` - print date, time from unix timestamp, and time difference between it and now
`ts -m 1577836800000` - works with milliseconds too
`ts 1577836800 1577838600` - print time difference between timestamps
`ts -m 1577836800000 1577838600000` - works with milliseconds too