SAS Coding Challenge

Versions used:
- Go 1.22.0
- VSCode 1.86.1



There are two (2) test data sets.
- One with comments for tracking what date-times are valid, invalid, and duplicate
- And one with the raw data that is read into the program (no comments) 

The valid date-times are spit out into a file that will be named "Output.txt"



Assumptions used in logic validation
- All lines of data will ONLY have a full ISO 8601 date-time format:     
    - YYYY-MM-DDThh:mm:ssTZD
- Data is read in as a string
- Time-zones only have 3 valid options “Z”, “+hh:mm”, “-hh:mm”. 
    - The minutes can only be in whole or 30 min increments per ISO 8601 standard. 
        - Valid: +08:00, +08:30, -07:00
        - Invalid: +08:45, or +08:23 is invalid
    - The range of acceptable time-zone designators is -12:30 to +12:30



Challenges experienced
- Learning Go language & syntax to create program (first program with Go!)
- Determine how to filter correct vs. incorrect formats
- Making sure test data accounts for all possible cases, and edge cases
    - Ensuring all time zone designators are accounted for
    - Ensuring years from 0000 - 9999 are valid
    - Time zone abbreviations (except for "Z") are invalid
        - Even having a "-" in -EST returns invalid



