# termsearcher
### Searching for right term in a clinical setting

This project is a quick hack to showcase a method for searching the right term in a healthcare codeset, subset or terminology. 
The intention is to help start a conversation around how a clinical documentation system can be integrated with a search system.

It can be plugged into any terminology or codeset like [SNOMED CT][], [ICD-10][], [CPT][].

Partially based on the [Search and Data Entry Guide][] provided by [IHTSDO][].

It provides often missed features in clinical term search interfaces like:
* On the fly spell checking
* Search modifiers in the text box (not look for a term, look for a combination etc)
* Auto-complete 
* Last searched - popular searches

### Here is what it looks like:

![alt text][logo]

It was built using Go and python and can be scaled to serve 50+ users on a single instance.

Please note that this code is not production ready.

[logo]: https://github.com/healthitinternals/termsearcher/blob/master/about/imgs/sct_srch_01.gif "Quick view on how it works"
[SNOMED CT]: https://www.snomed.org/snomed-ct
[ICD-10]: https://en.wikipedia.org/wiki/ICD-10
[CPT]: https://en.wikipedia.org/wiki/Current_Procedural_Terminology
[IHTSDO]: https://www.snomed.org
[Search and Data Entry Guide]: https://confluence.ihtsdotools.org/display/DOCSEARCH/Search+and+Data+Entry+Guide
