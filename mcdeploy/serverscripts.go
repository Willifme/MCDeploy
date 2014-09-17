package mcdeploy

// Will have todo

const Linuxscript = `#!/bin/sh 
	BINDIR=$(dirname "$(readlink -fn "$0")")
	cd "$BINDIR"
	java -Xmx1024M -jar craftbukkit.jar -o true`

const Macscript = `#!/bin/bash
	cd "$( dirname "$0" )"
	java -Xmx1024M -jar craftbukkit.jar -o true`

const Windowsscript = `java -Xmx1024M -jar craftbukkit.jar -o true
	PAUSE`
