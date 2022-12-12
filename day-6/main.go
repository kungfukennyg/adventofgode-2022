package main

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

const testInput string = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

func main() {
	type tc struct {
		in                    string
		expected, pt2Expected int
	}

	tcs := []tc{
		{
			in:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected:    7,
			pt2Expected: 19,
		},
		{
			in:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:    5,
			pt2Expected: 23,
		},
		{
			in:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:    10,
			pt2Expected: 29,
		},
		{
			in:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:    11,
			pt2Expected: 26,
		},
		{
			in:          `bjbffsfnsnppzpphvhjvjtjmjwjrjdjffwrfrvvrqrrqwrrqpqhqnnddvccrbbwcwbcbclclhlzlznntrrzffctcggzqgqtgqtgtrgttlhttgstgtsgsfsnsddsvdsvddrzrvrnvnrvnrnmrrvfvbfbnbmmtbbgpgtptjptpctpccmccgbccqbcqczqccdssfqfzzjgzzvcvgggrjggncgctgtjjpttqtrrvmrvmmzzfcffpgfgfqggbwgghcctllfhlffbbcffspfpcplpjjlwldwldlpddwzzlqzqfzzcwcmwwcddhgddstsnsjsvjsvvndndsnnclcvcpvpwpqwpqwqllcjcsjjbppfdftfrrpwwqtqtvtllrsrhhczzgllsbsvsttbsszzgwgjwjqwjwvvcmcbczbbgppghghgshhvmhmvvvflllllsrsqsbbfsshhrhddcncbnnrfnnbtbftfltftdtndtdtcdcbdcbchhtddhhvbvsbstsjsjppmpjpttmjjvtjvvdcctcbbbbqzzssfdsdrdwdpplspllphpnhnshnncwnnhlnhhvvpttfftwfwjfjzjsjddjldjlddvzzfzczrzjrzrhzhshslhsllmcllhhthrttfssnqnjnbbzfbbmpmlplhlttmzzwrwjrjppzmpmrprcprplrprmrvrvqrrnlnccswwvggvgcgddlrlrccqhhgvgnvvmlmmcqqhrqrwqwsswgssfjjgvjvppfffhjhrhfhhlphpmpzmpmlppcspcpgpnpddcmcllwvvqfqbbpbcchjjpzpzbbhbzhhrrzrhzhphqphhqvqttmhhbvbcbcmbbfgbbmnmncmnmvnvffgdfggvhvthvhqqdwwvmvlmljjjvsjswjwzwjjqmqrmqmlljqqjvqvtvlldjjfbbpbvbttqhhgrrcssfdsdccrncrncnmnvvmccnrrrtctjtsjttrmrhhvdhdrdjdpdhphsppnttnhhvdvvztvtftsfsrfrppsggfbgbbvbmmjgjrjttfnngttscclzcztctcltcllsmllggbfbhffjljhhbjjjgddzdvdvrvzrrgqgllnvllnccqjqhhwrhhtfttfpfsfvfqvqgvgpppjjhzzcszczbcbtblltlffdfwddfsspgprgpgjgdgsgcccrllbzbqzbzwbwwtmwwvpvlpljlsjssfnfqfzqffthtjtptvptvvmlvlzzdnznwnddjtjpptnpnccbsccjmjqqcfcdddwtdtftntzzpbbbhcbblglttrmmclcdchhzttjwtwztwzwzczfccmrmcrmrsrnsrnrsnrrfmmmcpmmjtjrrwssvdssrvrqvrvcrchrhfhpprwprwpwtttvrtvrvffdcfcssfrfbbvsslvvhphrrwggssjtjbttqtllnmlnljlmmtllfjjtjtpjttbwbswsnwnnjdjbblfbfmmblmmwnmnncjcwwhzzdffvppvmmtzzrhrchhtjhttcssvnsnvnhvvdgvvvgtvtqtgqqzrzbzqzgqgmqmtqqztqqgllchlclwwbwrrwprwprrlgrgwggjhggtgmtggjddfhdddnzztrtjrjprpwrwttqzqnzzfnfcnfcftflfttvbvpbpllrtltslttlddbzdbzzfjfrjfjcjhjssbcsczcnznlzzggdccllhbhpbhbjbwjbwjwtwffptpqqqmssrlrddwqwqdqfqjqwqtqhtsrgvsdtrjmhtgrwvwrbfqtgvjbwphbrszdcgtcpqqrcqtzvzjstzpbmwrqlrcsmlnqpsprsfnpqqsdzbfglcshtjpphmchdzggrcjwttwlzdffgpswzfjdcgzntgzsqvjdnwwwwmtjvpqjgmltmstzztpzflfdhbhhljgdmthnrdzhmtwcmpsjnlgwcvnwdbdrbhcsscgwtptzrmqcwdcmsssmqjpnvclhhbgsjrmqgvvqhgmhzmnpqtczqwmnpvvpdhdtdpdsrmzwtsfchzmdlggrnvwcfbfmrgffssdjjlwbdvjqqwddgmgfwvbzldccwnwrnltcrmblwqswjslnsjtfqvssdrdtpwptdtvtdwhgtgqnmhqfljjjsdgwmptbctjtpdhzmtshgdwnnhfjthmhdrqqqprrdwhvsfwfbvwtfvhgglfphzwjqffwcclbpmtcqzmtjmswscngtbmbdsvfzfbgwvhwlhtgdsnscnrssdqzvhmhplqppzrgdncfvvpnzhgnjrvcmhrqmzvzmdhhpjmrrnwfrhdgqdhvvstbldrgdcwbgjvcwhfrpbgrnvgcszhpbbgvqnvvrrcgprtsftjqtnrbqrzspmzpnchbrbbbpnjdllhnnbcfdsjjhhjcrvvtsgnnfvczvqgvbgphzzjcczsgtlvfrddzlvwdfhprnsvrnzdcfqsbfhcmnrgmrfqwcblbzgrpnvbtrqqnnfslfllfsmrrfsthfzgrwswdprzcswgrjcpzwfhzbhpmsjjtsgcqpnhtwpvpcbpttdpcftrqsbtgtspdhlvmphvnglsdntqfzcrwvvzsmjftjpgjglnnjhnpbhpcmwshdrfbczwtmtslcpmnpngvhlccvtwrsglrrfcmngshtjlnvrtqpfngrrtvhhvmnbwpjtwplfnmfqrbzzqzwchthjbbrpgdppmsjlbljrzqvhmsbrtwglgwnhmmdwpmvmjqqrhtmjjmcnbgtpbbqbnphzhwfhzddqbhqhtmwvffjdcfjmwjjdrqwsfzfrrwlmhndhdvrsdqtmccdqgppwpfrtcgzzwczfblvtjhhdjmdlldbtwthfdpjhpsgbcjjznmwdgnczbnfdfslhsjcjnthgsjlslbcvvqgqbstdwlqllpmmtqlqrbtnvphvwhbshhpdzfbclsqgdmhrnjmbwjzwzdtqswzgmnmwcqcmtpnjzzcrqftvnrdqghszchhmnvlvmcmblpcqnspsjthgqrdpbrzvwtfmnfdcfgwtlvpvwjdwzdvqvgdrcqvzwlbmcwmbsqfhzmwfqmjvgjtsprwbbsrldnwmvhrtsmcdsbftpvcsmpnlmlggmlrgjfvljmpfldftqhjmqqhwfpjtzrrhtlrmmstjphtmldslnnmfhnccrpgjmrbffcvgvmghhnpqhpvdqmdzqcjtcjplhlffwgslpsfzpwqsfpngscdlszlpctpqgdmvwfdfgpwrpltvlwrzrgjjjnjtrwctjsnbtpbfbgqzftmjhfrzrtmtnztlhwwgqnvmfnrshfcdswbqnlrqvtgjdzmqqcdgpwjlgfwnnnjsmmtfbvpwqvnjdjphclnvjntqlfwdppjgcvlcjmfdsbtgngcglmdgsgzwdvsqlvgwcrjtttgdmrlthhwhnnrvvrjgqzqmbcbmhdwmndhjstlmbtwjbgmlrqqcmqzjzcbbfrqrqlmvvgfrdtrwrpgcfsrnjdbfddwgwqrlpvfjgnjjnrlzpbtnjphlqzmdwnqhvblmwzvtnsvbcgqdpmgvchqmgjmbvrfwmmzlchhbqrfmvdffczcsjlhjrmmlmdztmltszrjlrgjwrlfwvlgqtqznnscbqgdzbvdnnjbfmcztjvbgbfmdhvrjgjcngtpzndpnpwwldlfrtqhwpfwphrgdzjvslnbpmrvjpcjpbbsmpwvzmldrspmrlbsptzfdngcscsllswzccjzlmbglsrthvbzznzpjdswhqncmrpnqhzggzzfvhlgqbvlmfsqglpphhswhjbpqnqfpzltmhndmmzclwfmlqztvrqdzfqjpdhttgshjwffdcchmvrwmblpzffbgwrgnqhhvvsvlwnzmmhjwrszpfdsncjwllrnzrsfjsrdgnrbjqlrvpmzbstlqdznhjgbslzmplnqprwqgddjlwzbtrmfsfdlggddqrccztjffvbnsmfdzdhrgsflffmmjtjlbtnfcwhwzdsnbbphbjlrfrddbpncjrtglsnrppdbznrbjqqzdswnhvssffhjzrwnmlvmwmljnhtsnplpjdjpqzbbmzzfcmpm`,
			expected:    1757,
			pt2Expected: 2950,
		},
	}

	// pt1
	for i, t := range tcs {
		actual, err := findMarkerPos(t.in, 4)
		if err != nil {
			panic(err)
		}
		if actual != t.expected {
			fmt.Printf("failed tc %d, expected %d, got %d\n", i, t.expected, actual)
		} else {

			fmt.Printf("tc %d passed\n", i)
		}
	}

	// pt2
	for i, t := range tcs {
		actual, err := findMarkerPos(t.in, 14)
		if err != nil {
			panic(err)
		}
		if actual != t.pt2Expected {
			fmt.Printf("failed tc %d, expected %d, got %d\n", i, t.pt2Expected, actual)
		} else {

			fmt.Printf("tc %d passed\n", i)
		}
	}
}

func findMarkerPos(in string, markerLen int) (int, error) {
	var buf bytes.Buffer
	buf.WriteString(in)

	lastFo := make([]string, markerLen)
	var read, markerCount int
OUTER:
	for {
		r, err := buf.ReadByte()
		if err != nil {
			return 0, errors.Wrapf(err, "failed to read byte")
		}
		rn := string(r)
		read++

		if markerCount >= markerLen {
			markerCount = 0
		}
		lastFo[markerCount] = rn
		markerCount++

		if read < markerLen {
			continue
		}

		for i, b := range lastFo {
			for j, c := range lastFo {
				if i == j {
					continue
				}

				if b == c {
					continue OUTER
				}
			}
		}

		return read, nil
	}

}
