package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestSimpleMerge(t *testing.T) {
	db1 := `<!DOCTYPE KEEPASSX_DATABASE>
<database>
 	<group>
 	  <title>Coding</title>
 	  <icon>1</icon>
 	  <group>
 	    <title>Sites</title>
 	    <icon>1</icon>
 	    <entry>
 	      <title>Github</title>
 	      <username>johndoe</username>
 	      <password>password1</password>
 	      <url>https://github.com/</url>
 	      <comment>Code and stuff</comment>
 	      <icon>1</icon>
 	      <creation>2013-03-17T15:22:22</creation>
 	      <lastaccess>2013-03-17T15:22:54</lastaccess>
 	      <lastmod>2013-03-17T15:22:54</lastmod>
 	      <expire>3479-01-01T00:00:00</expire>
 	    </entry>
 	  </group>
    <entry>
       <title>Slashdot</title>
       <username>johndoe</username>
       <password>password2</password>
       <url>https://www.slashdot.org/</url>
       <comment>News and Polls</comment>
       <icon>1</icon>
       <creation>2013-03-17T15:23:48</creation>
       <lastaccess>2013-03-17T15:24:12</lastaccess>
       <lastmod>2013-03-17T15:24:12</lastmod>
       <expire>3479-01-01T00:00:00</expire>
    </entry>
	</group>
</database>
`
	db2 := `<!DOCTYPE KEEPASSX_DATABASE>
<database>
 	<group>
 	  <title>Coding</title>
 	  <icon>1</icon>
 	  <group>
 	    <title>Sites</title>
 	    <icon>1</icon>
 	    <entry>
 	      <title>Github</title>
 	      <username>johndoe</username>
 	      <password>password3</password>
 	      <url>https://github.com/</url>
 	      <comment>Code and stuff</comment>
 	      <icon>1</icon>
 	      <creation>2013-03-17T15:22:22</creation>
 	      <lastaccess>2013-03-17T15:22:54</lastaccess>
 	      <lastmod>2013-03-15T15:22:54</lastmod>
 	      <expire>3479-01-01T00:00:00</expire>
 	    </entry>
 	  </group>
    <entry>
       <title>Slashdot</title>
       <username>johndoe</username>
       <password>password4</password>
       <url>https://www.slashdot.org/</url>
       <comment>News and Polls</comment>
       <icon>1</icon>
       <creation>2013-03-17T15:23:48</creation>
       <lastaccess>2013-03-17T15:24:12</lastaccess>
       <lastmod>2013-03-18T15:24:12</lastmod>
       <expire>3479-01-01T00:00:00</expire>
    </entry>
	</group>
</database>
`

	dbMerged := `<!DOCTYPE KEEPASSX_DATABASE>
<database>
  <group>
    <title>Coding</title>
    <icon>1</icon>
    <entry>
      <title>Slashdot</title>
      <username>johndoe</username>
      <password>password4</password>
      <url>https://www.slashdot.org/</url>
      <comment>News and Polls</comment>
      <icon>1</icon>
      <creation>2013-03-17T15:23:48</creation>
      <lastaccess>2013-03-17T15:24:12</lastaccess>
      <lastmod>2013-03-18T15:24:12</lastmod>
      <expire>3479-01-01T00:00:00</expire>
    </entry>
    <group>
      <title>Sites</title>
      <icon>1</icon>
      <entry>
        <title>Github</title>
        <username>johndoe</username>
        <password>password1</password>
        <url>https://github.com/</url>
        <comment>Code and stuff</comment>
        <icon>1</icon>
        <creation>2013-03-17T15:22:22</creation>
        <lastaccess>2013-03-17T15:22:54</lastaccess>
        <lastmod>2013-03-17T15:22:54</lastmod>
        <expire>3479-01-01T00:00:00</expire>
      </entry>
    </group>
  </group>
</database>`
	tempdir, _ := ioutil.TempDir(os.TempDir(), "kpxmerge-tests-")
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()
	ioutil.WriteFile(tempdir+"/db1.xml", []byte(db1), 0644)
	ioutil.WriteFile(tempdir+"/db2.xml", []byte(db2), 0644)

	files := globFiles(tempdir + "/*.xml")
	db0, err := NewDatabase(files[0])
	if err != nil {
		t.Fatalf("Failed to parse %s: %s", files[0], err)
	}
	for i := 1; i < len(files); i++ {
		dbN, err := NewDatabase(files[i])
		if err != nil {
			t.Fatalf("Failed to parse %s: %s. Skipping.", files[i], err)
		}
		db0.Merge(dbN)
	}
	// write out merged XML to STDOUT
	dbmStr := db0.String()
	dbmStr = trimLines(dbmStr)
	dbMerged = trimLines(dbMerged)
	if dbmStr != dbMerged {
		t.Errorf("Merged output invalid. Expected:\n%s\nGot:\n%s\n", dbMerged, dbmStr)
	}
}

func trimLines(str string) string {
	out := ""
	for _, line := range strings.Split(str, "\n") {
		out += strings.Trim(line, " ") + "\n"
	}
	return out
}
