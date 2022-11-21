package fuzzer

import (
	"log"

	"github.com/golang/mock/gomock"

	"github.com/favonia/cloudflare-ddns/internal/domainexp"
	"github.com/favonia/cloudflare-ddns/internal/mocks"
	"github.com/favonia/cloudflare-ddns/internal/pp"
)

type Reporter struct{}

func (r Reporter) Errorf(format string, args ...any) {
	log.Printf(format, args...)
}

func (r Reporter) Fatalf(format string, args ...any) {
	log.Panicf(format, args...)
}

func FuzzParseList(bytes []byte) {
	input := string(bytes)

	mockCtrl := gomock.NewController(Reporter{})
	mockPP := mocks.NewMockPP(mockCtrl)

	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: %v`, gomock.Any(), domainexp.ErrSingleAnd).AnyTimes()
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: %v`, gomock.Any(), domainexp.ErrSingleOr).AnyTimes()
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: unexpected token %q`, gomock.Any(), gomock.Any()).AnyTimes()        //nolint:lll
	mockPP.EXPECT().Warningf(pp.EmojiUserError, `Domain %q was added but it is ill-formed: %v`, gomock.Any(), gomock.Any()).AnyTimes() //nolint:lll
	mockPP.EXPECT().Warningf(pp.EmojiUserError, `Please insert a comma "," before %q`, gomock.Any()).AnyTimes()

	_, _ = domainexp.ParseList(mockPP, input)

	mockCtrl.Finish()
}

func FuzzParseExpression(bytes []byte) {
	input := string(bytes)

	mockCtrl := gomock.NewController(Reporter{})
	mockPP := mocks.NewMockPP(mockCtrl)

	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: %v`, gomock.Any(), domainexp.ErrSingleAnd).AnyTimes()
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: %v`, gomock.Any(), domainexp.ErrSingleOr).AnyTimes()
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: unexpected token %q`, gomock.Any(), gomock.Any()).AnyTimes()                  //nolint:lll
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: wanted %q; got %q`, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()      //nolint:lll
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: wanted %q; reached end of string`, gomock.Any(), gomock.Any()).AnyTimes()     //nolint:lll
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: wanted a boolean expression; got %q`, gomock.Any(), gomock.Any()).AnyTimes()  //nolint:lll
	mockPP.EXPECT().Errorf(pp.EmojiUserError, `Failed to parse %q: wanted a boolean expression; reached end of string`, gomock.Any()).AnyTimes() //nolint:lll
	mockPP.EXPECT().Warningf(pp.EmojiUserError, "Domain %q was added but it is ill-formed: %v", gomock.Any(), gomock.Any()).AnyTimes()           //nolint:lll
	mockPP.EXPECT().Warningf(pp.EmojiUserError, `Please insert a comma "," before %q`, gomock.Any()).AnyTimes()

	_, _ = domainexp.ParseExpression(mockPP, input)

	mockCtrl.Finish()
}
