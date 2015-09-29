#Onstoserver - A testpad for some terminology applications

#####A The processing of the problem diagnosis files for import into Sphinx
1. Seperate the semantic types from the fsn (Use ST3 - NP++ has issues)

* Regex match: \(([a-z])\w+\)"$
* Replace with:","$MATCH

1. Save the file as CSV (ProblemDiagnosis_SCT_V1.csv) and  create a new header at the top (semtyp). Header should now read: shortid,label,fsn,semtyp

* remove the extra space left at the end
* Find:"," 
* Replace:","

3. Import into a temporary excel sheet and remove the enclosing brackets from the semtyp column (can be done in ST3 but issues with brackets elsewhere)

* The semantic types that will show up are
* EVENT                                                 
* FINDING                                          
* DISORDER
* SITUATION

4. Remove the header, Add a new column on the left and start a counter from 1,2... till the end

5. Copy and paste into a new file in ST3 and save as tsv file - close the excel file without saving


#####B Load data into the index
1. Run indexer using: indexer.exe -c problemdiagnosis.conf sct_pd

2. To query the data, run the search service by: searchd -c problemdiagnosis.conf

#####The Semantic Types used are:
* EVENT
* FINDING
* DISORDER
* SITUATION (for situationonly take those with %history%)
* MORPHOLOGIC ABNORMALITY
* NAVIGATIONAL CONCEPT