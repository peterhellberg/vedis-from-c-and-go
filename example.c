/*
 * Compile this file together with the vedis engine source code to generate
 * the desired executable. For example:
 *  gcc test.c vedis.c -o test
 */
#include <stdio.h>
#include <stdlib.h>

#include "vedis.h"

/*
 * Display an error message and exit.
 */
static void Fatal(const char *zMsg)
{
  puts(zMsg);
  /* Shutdown the library */
  vedis_lib_shutdown();
  /* Exit immediately */
  exit(0);
}

/*
 * atexit() callback. Shutdown the Vedis library.
 */
void vedis_exit(void)
{
  vedis_lib_shutdown();
}

/* Vedis test */
int main(int argc, char **argv) {
  vedis *pStore;
  vedis_value *pResult;
  int rc;

  rc = vedis_open(&pStore, argc > 1 ? argv[1] : ":mem:");
  if( rc != VEDIS_OK ){
    Fatal("Vedis is not OK!");
  }

  /* Register the atexit() callback */
  atexit(vedis_exit);

  /* Exectute a command */
  rc = vedis_exec(pStore, "SET foo 'DATA FROM C'", -1);
  if( rc != VEDIS_OK ){
    Fatal("Vedis is not OK!");
  }

  /* Exectute another command */
  rc = vedis_exec(pStore, "GET foo", -1);
  if( rc != VEDIS_OK ){
    Fatal("Vedis is not OK!");
  }

  vedis_exec_result(pStore, &pResult);
  {
    const char *zResponse;
    /* Cast the vedis object to a string */
    zResponse = vedis_value_to_string(pResult,0);
    /* Output */
    printf(" foo: %s\n", zResponse);
  }

  /* Auto-commit and close the vedis handle */
  vedis_close(pStore);
  return 0;
}
