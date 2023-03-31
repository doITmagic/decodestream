package decodestream_test

import (
	"strings"
	"sync"

	"github.com/doITmagic/decodestream"
	gi "github.com/onsi/ginkgo/v2"
	gom "github.com/onsi/gomega"
)

var jsonString = `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}
`

var _ = gi.Describe("Decodestream", func() {

	stream := decodestream.NewJSONStream()

	gi.It("NewJSONStream", func() {
		gom.Expect(stream).NotTo(gom.BeNil())
	})

	gi.It("can return a wait channel", func() {
		gom.Expect(stream.Watch()).To(gom.BeAssignableToTypeOf(make(<-chan decodestream.Entry)))
	})

	gi.It("can start and return to wait channel", func() {
		results := make(map[int]interface{})
		var mx sync.Mutex
		go func() {
			defer gi.GinkgoRecover()
			i := 0
			for data := range stream.Watch() {
				if data.Error != nil {
					gom.Expect(data.Error).NotTo(gom.HaveOccurred())
				}
				mx.Lock()
				results[i] = data.Data
				mx.Unlock()
			}
		}()
		stream.Start(strings.NewReader(jsonString))
		_, closed := <-stream.Watch()
		if closed {
			gi.It("can receive 2 json objects", func() {
				gom.Expect(results).To(gom.HaveLen(2))
			})
			gi.It("first object has correct name", func() {
				gom.Expect(results[0]).To(gom.HaveKeyWithValue("AEAJM", gom.HaveKeyWithValue("name", "Ajman")))
			})
			gi.It("second object has correct name", func() {
				gom.Expect(results[1]).To(gom.HaveKeyWithValue("AEAJM", gom.HaveKeyWithValue("name", "Ajman")))
			})
		}
	})

})
