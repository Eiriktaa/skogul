// Package gen consists of auto-generated protobuf code for the junos
// streaming telemetry interface. This file provides the commands to
// (re)generate the files.
package gen

//go:generate /bin/bash -c "rm -f junos/telemetry/*pb.go; mkdir -p junos/telemetry"
//go:generate /bin/bash -c "tar xzf tar-balls/junos-telemetry-interface-21.3R1.tar.gz"
//go:generate /bin/bash -c "for a in junos-telemetry-interface/*.proto; do if echo $DOLLAR{a} | egrep -qv '/(gnmi|sr_|Gnmi)'; then protoc --gogo_out=junos/telemetry --gogo_opt=M=$DOLLAR{PWD}junos-telemetry-interface/ -Ijunos-telemetry-interface/ $DOLLAR{a}; else echo skipping $DOLLAR{a}; fi; done"
//go:generate /bin/bash -c "sed -i 's/^package.*/package telemetry/g' junos/telemetry/*.go"
