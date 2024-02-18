Role-Based Activity Guideline: 
Implement a program that reads a large list of date-time values from a file and writes to a separate file the list of unique, valid date-time values (no duplicates). A valid date-time value matches the following format (ISO 8601):
YYYY-MM-DDThh:mm:ssTZD

Where:
• YYYY = four-digit year
• MM = two-digit month (01 through 12)
• DD = two-digit day of month (01 through 31)
• hh = two digits of hour (00 through 23)
• mm = two digits of minute (00 through 59)
• ss = two digits of second (00 through 59)
• TZD = time zone designator (“Z” for GMT or +hh:mm or -hh:mm)

It is not necessary to perform semantic validation of the data-time value. In other words, the date-time value "9999-02-31T12:34:56+12:34" should be considered valid by your program even though February 31, 9999 is not a legitimate date.

* Guidelines and Requirements
* The program can be implemented in either Java, C, or Go (preferably).
* The validation logic may not use high-level library functions that perform format validation or regular expression parsing. However, low-level library functions may be used.
* Use development environment (IDE, OS platform) of your choice.
* Make sure your test data includes duplicate values, using both "Z" for GMT or +hh:mm or -hh:mm formats.
* Please commit the solution to a public git repository (eg, github), and let us know the location of the repository (preferably one day prior to your scheduled interview). Also be sure to commit your test data.
* Be prepared to present your solution and walk through how you designed, implemented, and tested the program.
* You will want to discuss assumptions you made and challenges you encountered.
* A more comprehensive solution is preferred, but the focus is on how you demonstrate your understanding of the solution and defend your design and implementation decisions.

Once everything is finalized, I will send you additional details and instructions, including dial-in information.