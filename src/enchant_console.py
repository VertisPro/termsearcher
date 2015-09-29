import os
import enchant
from enchant.checker import SpellChecker

#start using python -i enchant_console.py

# combine with the US dictionary
#pdict = enchant.DictWithPWL("en_US","wordlist.txt")

pdict = enchant.request_pwl_dict("wordlist.txt")
chkr = SpellChecker(pdict)

def SuggestSpelling(theword):
	return chkr.suggest(theword)

def CheckText(thetext):
	chkr.set_text(thetext)
	rsp = []
	for err in chkr:
		rsp.append(err.word)
	return rsp


'''
chkr.set_text("pelvi frecture")
for err in chkr:
	print "ERROR:", err.word

'''