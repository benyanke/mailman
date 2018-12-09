# Design notes


## High level

Design a three pane system initially, like outlook: folders, mail, and mail view.

start with reading imap
add smtp later


## Sections/modules needed

 * config reading (dotdir plus envs for some overrides like config path or debug mode)
 * tests
 * imap library (initial auth, fetch mailbox, fetch mail list, fetch individual mail info
 * gui layout
 * threaidng logic
 * config writing with gui panel for modifications which go back to config dir in human readable format
 * mail compose
 * smtp
 * signature
 * multiple accounts

## Nice to have future piceses:

 * CLI interface for checking mail and doing actions from scripts (cobra)
 * CalDAV contact integration w/ caldav server (with local cache of course)
 * caldav task integration
 * Calendar? Maybe too much?
