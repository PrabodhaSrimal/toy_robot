package toy_robot_test

import (
	. "."

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Robot", func() {

	Describe("Robot placement", func() {
        Context("initially", func() {
        	myrobo := Robot{-1, -1, NODIR}

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })
        })

        Context("when valid PLACE command executed", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(2, 1, "WEST")

            It("in valid state", func() {
                Expect(myrobo.Direction).To(Equal(WEST))
                Expect(myrobo.X).To(Equal(2))
                Expect(myrobo.Y).To(Equal(1))
            })
        })

        Context("when invalid PLACE command executed with wrong X", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(5, 1, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(-1, 1, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(20000, 1, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })
        })

        Context("when invalid PLACE command executed with wrong Y", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(1, 5, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, -1, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, 20000, "WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })
        })

        Context("when invalid PLACE command executed with wrong Direction", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(1, 1, " ")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, 1, "")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, 1, " WEST ")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, 1, "1WEST")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })

            myrobo.Place(1, 1, "2")

            It("in invalid state", func() {
                Expect(myrobo.Direction).To(Equal(NODIR))
                Expect(myrobo.X).To(Equal(-1))
                Expect(myrobo.Y).To(Equal(-1))
            })
        })

    })

    Describe("Robot movement", func() {
        Context("initialized and placed at 0,0,NORTH", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(0, 0, "NORTH")

            It("in 0,0,NORTH", func() {
                Expect(myrobo.Direction).To(Equal(NORTH))
                Expect(myrobo.X).To(Equal(0))
                Expect(myrobo.Y).To(Equal(0))
            })

        })

        Context("initialized and placed at 0,0,NORTH and move", func() {
        	myrobo := Robot{-1, -1, NODIR}
        	myrobo.Place(0, 0, "NORTH")
        	myrobo.Move()

            It("in 0,1,NORTH", func() {
                Expect(myrobo.Direction).To(Equal(NORTH))
                Expect(myrobo.X).To(Equal(0))
                Expect(myrobo.Y).To(Equal(1))
            })

        })
    })

})
