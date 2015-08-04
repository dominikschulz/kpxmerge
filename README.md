# KpxMerge

KpxMerge is a utility I've hacked to sync some diverged KeePassX databases.

Unfortunately KeePassX does not support merging key databases OOTB, but you can
export a nice XML file (KeePassX XML) from each KDB, merge them using this tool
and import them again. Then you can review the changes a save the merged KDB in place of
the original one(s).

## Usage

* Export from KeePassX. Go to Menu, Export, XML and select KeePassX XML
* Repeat the step above for each diverged KDB file
* Run kpxmerge /path/to/export1.xml /path/to/export2.xml ... /path/to/exportN.xml > /path/to/merged.xml
* Pay close attention to the log output and review any conflicts
* Import the merged.xml to KeyPassX. Go to Menu, Import, KeePassX XML
* Save the new KDB to a location of your choice

## WARNING
* DO NOT FORGET TO REMOVE THE XML FILES. THEY CONTAIN CLEARTEXT PASSWORDS
* USE WITH CAUTION. THIS TOOL TRIES TO BE SAVE BUT YOU MAY LOOSE ENTRIES
* CREATE A BACKUP OF YOUR KDB FILE BEFORE OVERWRITING IT

