\ API -> n4_api() declared functions
\ these functions control operations on the 
\ LCD display.
: append  ( count ... -- ) 1 API ; \ A.K.A. 'write'
: move_to ( col row   -- ) 2 API ;
: delete  ( up_to     -- ) 3 API ;

: $ 15 ; \ Shorthand for end of line (LCD 16x4 display)
   
\ NOTE: These names are just illustrative
: DRW ( row -- )
	0 move_to $ delete append
;

\		PORTAL   xxxxxx
\		REGIST.  xxxxxx
\		UNICAS   xxxxxx
\		COMUNICANDO WEB

VAR device
VAR tag_unique
VAR all_tag
VAR comm

: SC1 ( -- )

	device     @
	tag_unique @
	all_tag    @
	comm       @

	3 FOR
		I DRW
		50 DLY
	NXT

	0 DRW
;
