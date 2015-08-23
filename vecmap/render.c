//#include "_cgo_export.h"
extern goWriterFunc(void *ptr, const char *bytes, int len);

int cWriterFunc(void *ptr, const char *bytes, int len) {
  return goWriterFunc(ptr, bytes, len);
}

