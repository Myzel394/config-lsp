# /sshd_config/analyzer

This folder analyzes the config file and returns errors, warnings, and suggestions.

The analyzer is usually done in three steps:

1. Check if the overall structure is valid
2. Create indexes and check if they were created successfully
3. Do all other checks that require indexes

This is why there are three if statements to check for errors.
If the first step fails, we can't create any indexes and thus show these errors already.

