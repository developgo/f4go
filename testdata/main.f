C -----------------------------------------------------
C Tests 
C -----------------------------------------------------

        program MAIN
            ! begin of tests
            call testName("test_operations")
            call test_operations()

            call testName("test_pow")
            call test_pow()

            call testName("real_test_name")
            call real_test_name()

            call testName("test_subroutine")
            call test_subroutine()

            call testName("test_if")
            call test_if()

            call testName("test_do")
            call test_do()

            call testName("test_do_while")
            call test_do_while()

            call testName("test_array")
            call test_array()

            call testName("test_goto")
            call test_goto()

            call testName("test_function")
            call test_function()

            call testName("test_data")
            call test_data()

            call testName("test_matrix")
            call test_matrix()

            call testName("test_types")
            call test_types()

            call testName("test_concat")
            call test_concat()

            call testName("test_save")
            call test_save()

            call testName("test_complex")
            call test_complex()

            call testName("test_character")
            call test_character()

            call testName("test_matrix3")
            call test_matrix3()

            call testName("test_parameter")
            call test_parameter()

            call testName("test_write")
            call test_write()

            call testName("test_byte")
            call test_byte()

            call testName("test_vars")
            call test_vars()

            call testName("test_complex")
            call test_complex()

            call testName("test_assign")
            call test_assign()

            ! end of tests
        END

        subroutine testName(name)
            character(*) name
            write (*,FMT=45) name
            return
  45        FORMAT ('========== Test : ',A20,' ==========')
        end

        subroutine fail(name)
            character(*) name
            write (*,FMT=46) name
            STOP
  46        FORMAT ('***** FAIL : ',A)
        end

C -----------------------------------------------------
C -----------------------------------------------------
C -----------------------------------------------------
C -----------------------------------------------------

        subroutine real_test_name()
        end subroutine ! this is also test

C -----------------------------------------------------

        subroutine test_operations()
            integer i
            i = 12
            i = i + 12
            i = i - 20
            i = i *2
            write (*, FMT = 120) i ! TEST COMMENT
            write (*, '(I1,A2,I1)') i,'YY',i
            write (*, '()') ! TEST COMMENT
            return
  120 Format (' output ', I1 , ' integer I''C values ')
        end

C -----------------------------------------------------

        subroutine test_pow()
            REAL r,p
            Real real_1
            REAL*8 H(1)
            COMPLEX *16 C
            Intrinsic REAL
            C = (2.0,2.0)
C calculation
            H(1) = 3
            H(1) = H(1) ** H(1)
            r = -9.Q+4
            p = 1.45
            r = r / 5
            r = r + 1.4**2.1
            r = -(r + p**p)
            r = (r + p)**(p-0.6)
            real_1 = r/5
            real_1 = real_1 + 1.222**REAL(C)
            write (*, FMT = 125) r, real_1, H(1)
            return
  125 Format ('POW: ', F15.2 ,' , ', F14.2, ' , ' ,  F14.2)
        end

C -----------------------------------------------------

        subroutine test_subroutine()
            REAL a,b
            a = 5
            b = 6
            write (*, FMT = 131) 1,a
            write (*, FMT = 131) 2,b
            CALL ab(a,b)
            write (*, FMT = 131) 3,a
            write (*, FMT = 131) 4,b
            return
  131 Format (' output ', I2, ' ', F12.5 , ' real ')
        end

        subroutine ab(a,b)
            REAL a,b
            a = a + 5.12
            b = b + 8.4
            return
        end

C -----------------------------------------------------

        recursive subroutine test_if()
            integer i, J
            logical l
            l = .false.
            i = 5
            IF ( i .EQ. 5) write(*,*) "Operation  .EQ.    is Ok"
            IF ( i .GE. 5) write(*,*) "Operation  .GE.    is Ok"
            IF ( i .LE. 5) write(*,*) "Operation  .LE.    is Ok"
            IF ( i .GE. 4) write(*,*) "Operation  .GE.    is Ok"
            IF ( i .LE. 3) write(*,*) "Operation  .GE.    is Ok"
            IF ( i .NE. 3) write(*,*) "Operation  .NE.    is Ok"
            IF ( .true.  ) write(*,*) "Operation  .TRUE.  is Ok"
            IF (.NOT. l  ) write(*,*) "Operation .NOT. .false. is Ok"
            IF (.NOT. l  ) write(*,*) "Operation .NOT. .FALSE. is Ok"
            l = .TRUE.
            IF ( l       ) write(*,*) "Operation  .TRUE.  is Ok"

            IF (i .GE. 100) THEN
                STOP 
            ELSEIF (i .EQ. 5) THEN ! TEST COMMENT
                WRITE(*,*) "ELSEIF is Ok"
            END IF ! TEST COMMENT

            DO 100 J = 1,2
                IF (.NOT.l) THEN ! TEST COMMENT
                ELSE 
                    IF (J.GE.0) THEN 
                        WRITE(*,*) "Ok"
                    END IF
                END IF
  100       CONTINUE

            CALL ZD()

            RETURN
        end

        SUBROUTINE ZD()
			INTEGER IFORM
            DO 100 IFORM = 1, 2
      			IF ( IFORM .NE. 0 ) THEN
                    WRITE(*,'(I1)')IFORM
                END IF
  100       CONTINUE
        RETURN
        END
C -----------------------------------------------------

        subroutine test_do()
            integer IR, JR, HR, iterator
            integer ab_min
            JR = 1
            iterator = 1
            DO 140 IR = 1,2
                DO 130 HR = 1,2
                        WRITE (*,'(A8)') 'GOTO131'
  130           END DO
  140       END DO
            Do IR = 1,10,3
                write (*,FMT=142) IR
            end do
            Do IR = 1,3
                write (*,FMT=146) IR
            end do
            DO 143 IR = 1,3
                write (*,FMT=147) IR
  143 Continue
  144 Continue
            Do IR = 1,3
                write (*,FMT=148) IR
            end do

            if (ab_min(3,14) .EQ. 14) THEN
                call fail("test_do 1")
            end if

            Do IR = 1,ab_min(3,13)
                write (*,FMT=149) IR
            enddo
            DO 145, HR = 1,2
                DO 145 JR = 1,2
                    write(*,FMT=150) HR, JR
  145 Continue
            write (*,fmt=151)iterator
            iterator = iterator + 1
            IF ( iterator .LE. 3) THEN
                write(*,*) "iterator is less or equal 3"
                GO TO 144
            END IF
            call test_do2()
            return
  142 FORMAT ('Do with inc ', I2)
  146 FORMAT ('Do ', I2)
  147 FORMAT ('Do with continue ', I2)
  148 FORMAT ('Do with end do ', I2)
  149 FORMAT ('Do with enddo ', I2)
  150 FORMAT ('Double DO ', I2, I2)
  151 FORMAT (' iterator = ', I2)
         end

         integer function ab_min(a,b)
             integer a,b
             if (a .LE. b) then ! TEST COMMENT
                 ab_min = a
             else ! TEST COMMENT
                 ab_min = b
             end if ! TEST COMMENT
             return
         end function

        SUBROUTINE test_do2( )
            INTEGER I,J,N, VTSAV(2,2), A(2,2) ! TEST COMMENT
            N = 2 
            DO 140 J=1,N, 1 ! TEST COMMENT
               DO 130 I=1,N ! TEST COMMENT
                  A(I,J) = I+J*N ! TEST COMMENT
                  VTSAV(J,I) = A(I,J) ! TEST COMMENT
                  WRITE(*,'(I5)') VTSAV(J,I) ! TEST COMMENT
  130          END DO ! TEST COMMENT
  140       END DO ! TEST COMMENT
            RETURN ! TEST COMMENT
        END ! TEST COMMENT

C -----------------------------------------------------

        subroutine test_do_while()
            integer iterator
            iterator = 1
            Do while (iterator .Le. 3) ! TEST COMMENT
                write (*,180) iterator
                iterator = iterator + 1
            end do
            return
  180 FORMAT ('Do while ', I2)
        end 

C -----------------------------------------------------

        subroutine test_array()
            integer iterator(3),ir, summator
            external summator
            do ir = 1,3
                iterator(ir) = ir
            end do
            do ir = 1,3
                write(*,fmt = 210) iterator(ir)
            end do
            if (summator(iterator) .NE. 6) THEN
                call fail("test_array 1")
            end if
            return
  210 FORMAT ('vector ', I2)
        end 

        integer function summator(s)
            integer s(*), ir, sum
            sum = 0 
            do ir = 1,3
                sum = sum + s(ir)
            end do
            summator = sum
            return
        end function

C -----------------------------------------------------

        subroutine test_goto()
            integer t,m
            t = 0
            m = 0
  230       t = t + 1
  240       t = t + 5
            m = m + 1
            write(*, FMT = 251) t
            if ( m .GE. 5) THEN
                go to 250
            end if
            write(*, FMT = 252) m
            goto (230,240), m
  250       return
  251 FORMAT ('goto check t = ', I2)
  252 FORMAT ('goto check m = ', I2)
        end

C -----------------------------------------------------

        subroutine test_function()
            integer ai
            character bi*32
            logical l, function_changer
            external function_changer
            ai = 12
            bi = "rrr"
            write(*,fmt = 270) ai,bi
            l = function_changer(ai,bi)
            if ( l .NEQV. .TRUE.) THEN 
                call fail("test function in logical")
            end if
            write(6,fmt = 271) ai,bi
            return
  270 FORMAT('test function integer = ', I9, ' array = ' , A3)
  271 FORMAT('test function integer = ', I9, ' array = ' , A3)
        end subroutine

        logical function function_changer(a,b)
            integer a
            character b*32
            a = 34
            b = "www"
            function_changer = .TRUE.
            return
        end

C -----------------------------------------------------

        subroutine test_data
            INTEGER LOC12( 4 ), LOC21( 4 )
            INTEGER IPIVOT( 2, 2 )
            INTEGER IP( 5 )
            real*4  v
            integer r

            data v , r / 23.23 , 25 /
            DATA    LOC12 / 3, 4, 1, 2 / , LOC21 / 2, 1, 4, 3 /
            DATA    IPIVOT / 1, 2, 3, 4 /

            INTEGER     LV
            PARAMETER ( LV = 2 )
            INTEGER MM( LV, 4 ), J, NN
            DATA    ( MM( 1, J ), J = 1, 4 ), NN  /494,322,2508,2549,42/
            DATA    ( IP(    J ), J = 1, 4 ) / 12,14,16,18/
            DATA    IP(5) / 123 /
            DATA    MM(2,3) / 56/

            INTEGER KTYPE(21)
            DATA    KTYPE / 1, 2, 3, 5*4, 4*6, 6*6, 3*9 /

            CHARACTER CA(4)
            DATA    CA / 'A', 'B', 'C', 'D'/
            
            logical XSWPIV(4)
            DATA    XSWPIV / .FALSE., .FALSE., .TRUE., .TRUE. /

            Logical L0(-1:1)
            DATA    L0 / .FALSE.,.FALSE., .FALSE./
            Logical L1(0:1, 0:1)
            DATA    L1 / .FALSE.,.FALSE.,.FALSE.,.FALSE. /

            DATA CHEPS/2.22D-16/
            IF (CHEPS - 2.22D-16 .LT. 0) THEN
                CALL FAIL("TEST_DATA: CHEPS")
            END IF

            if ( r .NE. 25) then 
                call fail("test_data 1")
            end if
            if(.NOT. ( 23.0 .LE. v .AND. v .LE. 23.5 )) then
                call fail("test_data 2")
            end if 
            if ( LOC12(1) .NE. 3) call fail("test_data 3")
            if ( LOC12(2) .NE. 4) call fail("test_data 4")
            if ( LOC12(3) .NE. 1) call fail("test_data 5")
            if ( LOC12(4) .NE. 2) call fail("test_data 6")

            if ( LOC21(1) .NE. 2) call fail("test_data 7")
            if ( LOC21(2) .NE. 1) call fail("test_data 8")
            if ( LOC21(3) .NE. 4) call fail("test_data 9")
            if ( LOC21(4) .NE. 3) call fail("test_data 10")

            if (IPIVOT(1,1) .NE. 1) call fail("test_data matrix 1")
            if (IPIVOT(2,1) .NE. 2) call fail("test_data matrix 2")
            if (IPIVOT(1,2) .NE. 3) call fail("test_data matrix 3")
            if (IPIVOT(2,2) .NE. 4) call fail("test_data matrix 4")
            
            if (IP(1) .NE. 12 ) call fail("test_data IP 1")
            if (IP(2) .NE. 14 ) call fail("test_data IP 2")
            if (IP(3) .NE. 16 ) call fail("test_data IP 3")
            if (IP(4) .NE. 18 ) call fail("test_data IP 4")
            if (IP(5) .NE.123 ) call fail("test_data IP 5")

            if (MM(1,1) .NE. 494 ) call fail("test_data MM 1 1")
            if (MM(1,2) .NE. 322 ) call fail("test_data MM 1 2")
            if (MM(1,3) .NE. 2508) call fail("test_data MM 1 3")
            if (MM(1,4) .NE. 2549) call fail("test_data MM 1 4")
            if (MM(2,3) .NE. 56  ) call fail("test_data MM 2 3")

            if (NN      .NE. 42  ) call fail("test_data NN")

            if (XSWPIV(1)) call fail("test_data 1")
            if (XSWPIV(2)) call fail("test_data 2")
            if (.NOT.XSWPIV(3)) call fail("test_data 3")
            if (.NOT.XSWPIV(4)) call fail("test_data 4")

            IF (L0(-1)) call fail("test_data L0 -1")
            IF (L0(0)) call fail("test_data L0 0")
            IF (L0(1)) call fail("test_data L0 1")

            IF (L1(0,0)) call fail("test_data L1(0,0)")
            IF (L1(1,0)) call fail("test_data L1(1,0)")
            IF (L1(0,1)) call fail("test_data L1(0,1)")
            IF (L1(1,1)) call fail("test_data L1(1,1)")

            DO J = 1,21
                Write(*,'(I5)') KTYPE(J)
            END DO
            DO J = 1,4
                Write(*,'(A1)') CA(J)
            END DO

            write(*,'(A2)') "ok"
            write(*,'(I2)') LV

        end subroutine

C -----------------------------------------------------

        subroutine test_matrix
            integer M(3,2),I,J
            do I = 1,3
                do J = 1,2
                    M(I,J) = I*8+J+(I-J)*5
                end do
            end do
            do I = 1,3
                do J = 1,2
                    write(*,fmt=330) I, J, M(I,J)
                end do
            end do
            call matrix_changer(M,3,2)
            do I = 1,3
                do J = 1,2
                    write(*,fmt=331) I, J, M(I,J)
                end do
            end do
            return
  330 format('Matrix (',I1,',',I1,') = ', I2 )
  331 format('Matrix*(',I1,',',I1,') = ', I2 )
        end

        subroutine matrix_changer(M,IN,JN)
            integer M(IN,JN), IN, JN, I, J
            do I = 1,IN
                do J = 1,JN
                    M(I,J) = M(I,J) + 2*(I+J)
                end do
            end do
            return
        end subroutine

C -----------------------------------------------------

        subroutine test_types
            REAL    R1
            REAL *4 R4
            REAL *8 R8
            DOUBLE precision DP
            INTEGER    I1
            INTEGER *2 I2
            INTEGER *4 I4
            INTEGER *8 I8
            R1 = 45.1
            R4 = 45.1
            R8 = 45.1
            DP = 45.1
            I1 = 12
            I2 = 12
            I4 = 12
            I8 = 12
            if ( 45.0 .LE. R1 .AND. R1 .LE. 45.2) THEN
                write(*,*)'R1 ... ok'
            end if
            if ( 45.0 .LE. R4 .AND. R4 .LE. 45.2) THEN
                write(*,*)'R4 ... ok'
            end if
            if ( 45.0 .LE. R8 .AND. R8 .LE. 45.2) THEN
                write(*,*)'R8 ... ok'
            end if
            if ( 45.0 .LE. DP .AND. DP .LE. 45.2) THEN
                write(*,*)'DP ... ok'
            end if
            if ( I1 .Eq. 12) write(*,*)'I1 ... ok'
            if ( I2 .Eq. 12) write(*,*)'I2 ... ok'
            if ( I4 .Eq. 12) write(*,*)'I4 ... ok'
            if ( I8 .Eq. 12) write(*,*)'I8 ... ok'
            return
        end subroutine

C -----------------------------------------------------

        subroutine test_concat
            CHARACTER*1 a,b
            CHARACTER*2 c
            a = 'a'
            b = 'b'
            c = a // b
            write(*,'(A1)')a
            write(*,'(A1)')b
            write(*,*)c
        end

C -----------------------------------------------------

        subroutine test_save
            integer iter 
            iter = 1
            call save_sub(iter)
            call save_sub(iter)
        end

        subroutine save_sub(iter)
            integer r,iter
            Save r
            DATA r /0/
            iter = iter + 1
            write(*,'(I2,I2)') iter , r
        end

C -----------------------------------------------------

        subroutine test_complex
            complex ONE
            complex c1
            complex*8 c2,c3
            COMPLEX*16 zc
            double complex db1, db2
            PARAMETER (ONE= (1.9E+0,2.3E+0))
            REAL    * 8   R
            COMPLEX * 16  CR
            COMPLEX * 16  CRA(1)
            Intrinsic real , aimag

            c1 = (1.1221,2.2)
            c2 = (3.23,-5.666)
            c1 = c1 + c2
            write(*,fmt = 500) real(ONE), aimag(ONE)
            WRITE(*,*)'CONJG'
            write(*,fmt = 500) real(c1),  aimag(c1)
            c1 = conjg(c1)
            write(*,fmt = 500) real(c1),  aimag(c1)
            zc = (-2.333,5.666)
            WRITE(*,*)'DCONJG'
            write(*,fmt = 500) real(zc),  aimag(zc)
            zc = dconjg(zc)
            write(*,fmt = 500) real(zc),  aimag(zc)

            c3 = c1 - ONE
            call comp_par(c3)

            db1 = (2.333,3.444)
            db2 = (1.222,0.111)
            write(*,fmt = 500) real(db1), aimag(db1)
            db1 = db2 * db1
            write(*,fmt = 500) real(db1), aimag(db1)
            write(*,fmt = 500) real(db2), aimag(db2)

            write(*,*) "==== REAL * COMPLEX ===="
            CR = (0.12,0.34)
            R = 12.34
            CR = R * CR
            write(*,fmt = 500) real(cr), aimag(cr)
            CR = CR * R
            write(*,fmt = 500) real(CR), aimag(CR)
            CR = (1.23,2.34)
            CR = DBLE(CR) * CR
            write(*,fmt = 500) real(CR), aimag(CR)
            CR = CR * DBLE(CR)
            write(*,fmt = 500) real(CR), aimag(CR)
            CRA(1) = (3.45,4.56)
            CR = CRA(1) * CR
            write(*,fmt = 500) real(CR), aimag(CR)
            CR = CR * CRA(1)
            write(*,fmt = 500) real(CR), aimag(CR)

            return
  500  FORMAT('>',F15.2,' + ',F15.2,'<')
       end

        subroutine comp_par(c)
            complex c
            write(*,fmt = 500) real(c), aimag(c)
            return
  500  FORMAT('>',F4.2,' + ',F4.2,'<')
       end


C -----------------------------------------------------

       SUBROUTINE test_character
           CHARACTER*5 CH(2)
           CHARACTER*6 CT(2)
           INTEGER NW, NS
           PARAMETER (NW = 6)
           PARAMETER (NS = 2)
           DATA CT(1) /'123456'/
           DATA CT(2) /'ABCDFE'/
           CHARACTER*6 CS(NS)
           DATA CS /'123456', 'ABCDEF' /
           WRITE(*, FMT=531 ) CT(1)
           WRITE(NW, FMT=531 ) CT(2)
           CH(1) = 'qwe'
           CH(2) = 'asd'
           WRITE(*, FMT=530 ) CH(1)
           WRITE(6, FMT=530 ) CH(2)
           WRITE(*, FMT=531 ) CS(1)
           WRITE(*, FMT=531 ) CS(2)
           WRITE(*, '(I2,I2)') NW, NS
           RETURN
  530      FORMAT('-->',A3)
  531      FORMAT('++>',A6)
       END SUBROUTINE

C -----------------------------------------------------

       SUBROUTINE test_matrix3
           INTEGER V
           PARAMETER(V = 3)
           INTEGER IR(2,V,4), I,J,K
           COMPLEX  CV(3,2,1)
           INTEGER IM3(2,3,4)
           DATA   ((CV(I,J,1),I=1,3),J=1,2)/(0.1E0,0.1E0),
     +            (1.0E0,2.0E0), (2.0E0,3.0E0), (3.0E0,4.0E0),
     +            (5.0E0,6.0E0), (6.0E0,7.0E0)/
           DATA IM3/1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6,7,8,9,0,
     +              1,2,3,4/
           DO I = 1, 2
                DO J = 1, 3
                    DO K = 1, 4
                        IR(I,J,K) = I*27+J*13+K
                    END DO
                END DO
           END DO
           DO I = 1, 2
                DO J = 1, 3
                    DO K = 1, 4
                        WRITE(*,FMT=565)I,J,K, IR(I,J,K)
                    END DO
                END DO
           END DO
           DO I = 1, 2
                DO J = 1, 3
                    DO K = 1, 4
                        WRITE(*,FMT=567)I,J,K, IM3(I,J,K)
                    END DO
                END DO
           END DO
           DO I = 1, 3
                DO J = 1, 2
                    DO K = 1, 1
                        WRITE(*,FMT=566)I,J,K, REAL(CV(I,J,K)),
     +                     AIMAG(CV(I,J,K))
                    END DO
                END DO
           END DO
           WRITE(*, '(I2)') V
           RETURN
  565      FORMAT('IR(',I1,',',I1,',',I1,')=',I3)
  566      FORMAT('CV(',I1,',',I1,',',I1,')=',F5.2,'::',F5.2)
  567      FORMAT('IM3(',I1,',',I1,',',I1,')=',I4)
       END SUBROUTINE

C -----------------------------------------------------

        SUBROUTINE test_parameter
            REAL      ZERO, ONE
            PARAMETER ( ZERO = 0.0E+0, ONE = 1.0E+0 )
            WRITE(*,FMT=570) ZERO
            WRITE(*,FMT=570) ONE
            RETURN
  570       FORMAT('PARAMETER : ', F10.4, /10x, 'values')
        END SUBROUTINE

C -----------------------------------------------------

        SUBROUTINE test_write
            INTEGER OUT
            REAL VALUE
            ! CHARACTER*512 IN_CH
            INTEGER       IN_I
            REAL          IN_R1
            REAL          IN_A(3)
            ! INTEGER       NN
            ! INTEGER       NVAL(60)
            INTEGER       A,I
            INTEGER       NAI
            PARAMETER     (NAI = 5)
            INTEGER       AI(NAI)
            ! LOGICAL       L1, L2
            ! L1 = .TRUE.
            ! L2 = .FALSE.

            ! NN = 20
            ! DO I = 1,60
            !     NVAL(I) = I*5
            ! END DO
            ! WRITE( *, FMT = 620 )'N   ', ( NVAL( I ), I = 1, NN )
            ! WRITE( *, '(I5)'    ) NN 
C -----------------
            A = 3
            OUT   = 6
            VALUE = 12.34
            WRITE(OUT,FMT=615) VALUE
            OPEN(UNIT=2,FILE = "./testdata/text")
            OUT = 2
C
C           INTEGER
C
            READ (UNIT = OUT, FMT = 617 ) IN_I
            WRITE(*  ,'(I7)' ) IN_I
            READ (OUT,'(I7)' ) IN_I
            WRITE(*  ,'(I7)' ) IN_I
            READ (OUT,'(I7)' ) IN_I
            WRITE(*  ,'(I7)' ) IN_I
            READ (OUT,'(I7)' ) IN_I
            WRITE(*  ,'(I7)' ) IN_I
            READ (OUT,'(I7)' ) IN_I
            WRITE(*  ,'(I7)' ) IN_I
            READ (OUT, FMT=* ) IN_I
            WRITE(*  ,'(I7)' ) IN_I

            ! TODO 
  !     DO 100 I = 1, N
  !        READ( NIN, FMT = * )( A( I, J ), J = 1, N )
  ! 100 CONTINUE

            READ( OUT, FMT = * )( AI ( I ), I = 1, NAI )
            DO I = 1, NAI
                WRITE(*,'(I2)') AI(I)
            END DO
C
C           REAL
C
            READ (OUT,'(F6.2)') IN_R1
            WRITE(UNIT= 6 ,FMT = '(F8.2)') IN_R1
            ! WRITE(UNIT= * ,FMT = '(D8.2)') IN_R1
            ! WRITE(UNIT= 6 ,FMT = '(E8.2)') IN_R1
C
C           REAL ARRAY
C
            READ (OUT, FMT = *  ) (IN_A(I), I=1,A)
            WRITE( * , '(F8.2)' )  IN_A(1)
            WRITE( * , '(F8.2)' ) (IN_A(I), I=1,A)
C
C           CHARACTER
C
            ! READ (OUT,'(A80)') IN_CH
            ! WRITE(*  ,'(A80)') IN_CH
            CLOSE(2)
C
C           LOGICAL
C
            ! WRITE(*,'(L1)') L1
            ! WRITE(*,'(L1)') L2
            ! WRITE(*,'(L2)') L1
            ! WRITE(*,'(L2)') L2

            WRITE(*,FMT=618)
            WRITE(*,FMT=619)
            RETURN
  615       FORMAT ('TEST_WRITE :',F6.3)
  617       FORMAT (I7)
  618       FORMAT ( ' !! Invalid input value: ', '=', '; must be <=' )
  619       FORMAT ( ' UPLO=''', ''', DIAG=''', ''', N=', ', NB=' , ',
     $      test(', ')= ')
  ! 620       FORMAT( 4X, A4, ':  ', 10I6, / 11X, 10I6 )
        END SUBROUTINE

C -----------------------------------------------------

        SUBROUTINE test_byte
            CHARACTER(1) SRNAME_ARRAY(3)
            DATA SRNAME_ARRAY / 'q' , 'w' , 'e' /
            call test_byte_func(SRNAME_ARRAY,3)
        END SUBROUTINE

        SUBROUTINE test_byte_func(SRN,L)
            INTEGER L,I
            CHARACTER(1) SrN(L)
            INTRINSIC MIN, LEN
            CHARACTER*32 S
            DO I = 1, MIN(L, LEN(S))
                S(I:I) = SRn(I)
                WRITE(*,'(A1,A1)') sRN(I), S(I:I)
                END DO
            RETURN
        END SUBROUTINE

C -----------------------------------------------------

        SUBROUTINE test_vars
            INTEGER TTT
            TtT = 1
            WRITE(*,'(I2)') tTT
        END

C -----------------------------------------------------

        SUBROUTINE test_assign
            INTEGER TMP , TMPA 
            INTEGER TMP2, TMPA2
            TMP  = 10
            TMPA = 865
            IF ( TMP .EQ. 10) ASSIGN 860 TO TMPA
            TMP = TMP + 10
  860       TMP = TMP + 25
  865       TMP = TMP + 45
            WRITE(*,'(I2)') TMP

            TMP2  = 15
            TMPA2 = 875
            IF ( TMP .NE. 10) ASSIGN 870 TO TMPA2
            TMP2 = TMP2 + 10
  870       TMP2 = TMP2 + 25
  875       TMP2 = TMP2 + 45
            WRITE(*,'(I2)') TMP2
        END

C -----------------------------------------------------
