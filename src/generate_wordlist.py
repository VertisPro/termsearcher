# Generates unique words from a set of labels
# Export the label from the refsets.problemdiagnosissct to a csv file
# This script generates a wordlist.txt which will need to be filtered and sorted out further
# The filtering is currently done via excel but a good tool may be very useful.

import string

with open('wordlist.csv', 'rb') as f:
    content = f.read()
    content = content.translate(None, string.punctuation).lower()
    words = content.split()

with open("wordlist.txt", "w") as text_file:
    for word in words:
  		text_file.write("%s\n" % word)

