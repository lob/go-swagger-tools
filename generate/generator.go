//go:generate go-bindata -o "$SPEC_DIR/bindata.go" -prefix $SPEC_DIR -ignore .+go$ -nomemcopy -pkg $PKG_NAME $SPEC_DIR
//go:generate swagger generate spec -o $SPEC_DIR/swagger.json -b $START_FILE --scan-models
//go:generate go-bindata -o "$SPEC_DIR/bindata.go" -prefix $SPEC_DIR -ignore .+go$ -nomemcopy -pkg $PKG_NAME $SPEC_DIR

// Package generate provides go:generate shorthand for building swagger.json
//
package generate
