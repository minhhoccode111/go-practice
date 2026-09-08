#!/bin/bash

if command -v fd &>/dev/null; then
  FD_CMD=(fd -t f)
elif command -v fdfind &>/dev/null; then
  FD_CMD=(fdfind -t f)
elif command -v find &>/dev/null; then
  FD_CMD=(find . -type f)
else
  echo "Error: no file finder available" >&2
  exit 1
fi

"${FD_CMD[@]}" -x file --mime-type {} \; \
  | grep -E 'application/x-executable|application/x-sharedlib' \
  | cut -d: -f1 \
  | xargs -r rm -v
