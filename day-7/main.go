package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const testInput string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

const test2 string = `$ cd /
$ ls
dir a
$ cd a
$ ls
dir a
2 a.txt
$ cd a
$ ls
99999 a.txt`

func main() {
	tr, err := parseTerminal(pt1Input)
	if err != nil {
		panic(err)
	}

	totalSpace := 70_000_000
	updateSize := 30_000_000
	// 6400111
	fmt.Printf("we have total space %d\n", totalSpace)
	tr.findWithinSize(nil, 30_000_000)
	freeSpace := Abs(totalSpace - tr.size - updateSize)
	fmt.Printf("disk space consumed is %d\n", tr.size)
	fmt.Printf("therefore we have %d bytes to free\n", freeSpace)
	fmt.Printf("out of %d needed for update\n", updateSize)

	fmt.Println(tr.Print(1))
	ch := make(chan int)
	go tr.findWithinSize(ch, freeSpace)
	smallest := 99999999999
OUTER:
	for {
		select {
		case size := <-ch:
			if smallest > size {
				fmt.Printf("updating smallest size to %d\n", size)
				smallest = size
			}
		case <-time.After(1 * time.Second):
			break OUTER
		}
	}
	fmt.Printf("smallest: %d\n", smallest)
	fmt.Printf("root: %d\n", tr.size)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseTerminal(in string) (*tree, error) {
	root := New("/")

OUT:
	for _, line := range strings.Split(in, "\n") {
		fmt.Printf("%s\n", line)
		isCmd := strings.Index(line, "$") != -1
		if isCmd {
			line = line[2:]
			cmd, err := parseCommand(line)
			if err != nil {
				return root, errors.Wrapf(err, "failed to parse command '%s'", line)
			}
			switch cmd.typ {
			case commandChangeDir:
				if cmd.subj == ".." {
					if root.parent != nil {
						root = root.parent
					}
					fmt.Printf("%s\n", root.getPath())
					continue OUT
				}

				root = root.findChildDir(cmd.subj)
				fmt.Printf("%s\n", root.getPath())
			case commandList:
				// no-op to do, wait for next iteration
			}
		} else {
			// output from prev command, most likely ls
			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				return root, errors.New(fmt.Sprintf("expected at least 2 parts for line '%s'", line))
			}
			if parts[0] == "dir" {
				root = root.insertDir(parts[1])
			} else {
				size, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					return root, err
				}

				name := parts[1]
				root = root.insertFile(name, int(size))
			}
		}
	}

	p := root
	for {
		fmt.Printf("looking for parent, at %v\n", p)
		if p.parent == nil {
			fmt.Printf("found parent %+v\n", p)
			return p, nil
		}

		p = p.parent
	}
}

type leaf struct {
	name string
	size int
}

type tree struct {
	parent *tree
	dir    string
	leaves []leaf
	dirs   []*tree
	size   int
}

func (t *tree) findWithinSize(ch chan<- int, max int) int {
	size := 0
	if t == nil {
		return size
	}

	for _, val := range t.leaves {
		size += val.size
	}

	for _, dir := range t.dirs {
		size += dir.findWithinSize(ch, max)
	}

	if size > 0 && size >= max {
		if ch != nil {
			ch <- size
		}
	}
	t.size = size
	return size
}

func (t *tree) Print(indent int) string {
	ret := ""
	if t.dir == "/" {
		ret += fmt.Sprintf("- / (dir)\n")
	}
	print := func(in string) string {
		r := ""
		for i := 0; i < indent; i++ {
			r += "  "
		}
		return r + in
	}

	for _, dir := range t.dirs {
		ret += print(fmt.Sprintf("- %s (dir)\n", dir.dir))
		ret += dir.Print(indent + 1)
	}

	for _, leaf := range t.leaves {
		ret += print(fmt.Sprintf("- %s (file, size=%d\n", leaf.name, leaf.size))
	}
	return ret
}

func (t *tree) getPath() string {
	str := t.dir
	if t.parent == nil {
		return str
	}

	return path.Join(t.parent.getPath(), str)
}

func New(root string) *tree {
	tr := &tree{dir: "/"}
	return tr
}

func (t *tree) findChildDir(name string) *tree {
	if t == nil {
		return nil
	}

	for _, t := range t.dirs {
		if t.dir == name {
			return t
		}
	}

	fmt.Printf("failed to find child dir %s in %+v\n", name, t)
	return t
}

func (t *tree) insertDir(name string) *tree {
	if t == nil {
		return nil
	}

	if t.dirs == nil {
		t.dirs = []*tree{}
	}

	t.dirs = append(t.dirs, &tree{dir: name, parent: t})
	return t
}

func (t *tree) insertFile(name string, size int) *tree {
	if t == nil {
		return nil
	}

	if t.leaves == nil {
		t.leaves = []leaf{}
	}

	lf := leaf{
		name: name,
		size: size,
	}
	t.leaves = append(t.leaves, lf)
	return t
}

type kind string

const (
	commandChangeDir kind = "cd"
	commandList      kind = "ls"
)

func kindFromStr(in string) (kind, error) {
	switch in {
	case "cd":
		return commandChangeDir, nil
	case "ls":
		return commandList, nil
	default:
		return "", errors.New(fmt.Sprintf("unrecognized command %s", in))
	}
}

type command struct {
	typ  kind
	subj string
}

func cmdFromStr(parts []string) (command, error) {
	typ, err := kindFromStr(parts[0])
	if err != nil {
		return command{}, err
	}
	cmd := command{
		typ: typ,
	}
	if len(parts) > 1 && typ != commandList {
		cmd.subj = parts[1]
	}
	return cmd, nil
}

func parseCommand(line string) (command, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 1 {
		return command{}, errors.New(fmt.Sprintf("command of len %d too short", len(parts)))
	}

	return cmdFromStr(parts)
}

const pt1Input string = `$ cd /
$ ls
282959 btm
dir fmfnpm
dir gwlwp
dir hchp
275929 hmbbjbf
dir nsphznf
dir phschqg
193293 rhpwvff
dir spfwthmd
dir wchdqb
dir zlpmfh
191479 zlpmfh.gpt
$ cd fmfnpm
$ ls
dir fgtqvq
194704 fwdvgnqp.fsm
48823 fwdwq.tsq
224991 mtjngt
79386 rdsgpfjb.sfn
dir rvnwwfq
dir wrzcjwc
dir zlpmfh
$ cd fgtqvq
$ ls
293783 rjc.ncl
324635 wdjrhw
$ cd ..
$ cd rvnwwfq
$ ls
76914 btm
$ cd ..
$ cd wrzcjwc
$ ls
dir fwdwq
2159 fzb.tjs
dir lddhdslh
dir mjp
284475 vclnlds
196284 zjtftd
$ cd fwdwq
$ ls
120795 jqnl.hzj
$ cd ..
$ cd lddhdslh
$ ls
293030 fzb
dir gzj
$ cd gzj
$ ls
dir qzgsswr
$ cd qzgsswr
$ ls
33681 qzgsswr.wmv
121649 sbjbw.shv
$ cd ..
$ cd ..
$ cd ..
$ cd mjp
$ ls
289491 btm
169221 jqnl.hzj
$ cd ..
$ cd ..
$ cd zlpmfh
$ ls
189296 ldgpvnh
$ cd ..
$ cd ..
$ cd gwlwp
$ ls
dir dwcrnbj
dir fmfnpm
dir fwdwq
dir hzpsts
dir hzrq
dir jzwpjtf
dir lmmpmghg
dir mnw
dir qzgsswr
dir zlpmfh
$ cd dwcrnbj
$ ls
182989 btm
145822 fwdvgnqp.fsm
dir jbtfslcn
dir lgbglc
293584 mfstl.hhp
dir sbffqq
dir zhvn
$ cd jbtfslcn
$ ls
195255 sbjbw.shv
$ cd ..
$ cd lgbglc
$ ls
261423 fmfnpm.rqh
323530 fzb.lmm
314800 hbl.blm
173052 lwpglgt
dir qdpjss
184808 zlpmfh.chl
$ cd qdpjss
$ ls
83714 djsrhr.vch
68191 lqljcq.sdv
$ cd ..
$ cd ..
$ cd sbffqq
$ ls
29842 fwdvgnqp.fsm
$ cd ..
$ cd zhvn
$ ls
dir clpcg
dir gswvch
dir lgmfhnq
184407 rlvprbs
235779 rnh.dlv
99852 ttjtnj.gjs
$ cd clpcg
$ ls
97780 fwdvgnqp.fsm
201515 rth.rhm
$ cd ..
$ cd gswvch
$ ls
dir fzb
dir mdrlrtl
$ cd fzb
$ ls
dir qzgsswr
$ cd qzgsswr
$ ls
198805 jqnl.hzj
$ cd ..
$ cd ..
$ cd mdrlrtl
$ ls
154648 spm.wvf
$ cd ..
$ cd ..
$ cd lgmfhnq
$ ls
303154 mjzw.tbq
$ cd ..
$ cd ..
$ cd ..
$ cd fmfnpm
$ ls
134756 bdmdfvs.llm
228983 brfrv
dir fzb
52921 mnfsltsj
185748 sbjbw.shv
dir svtswb
$ cd fzb
$ ls
dir fzb
$ cd fzb
$ ls
dir gwtslbr
$ cd gwtslbr
$ ls
224479 fmfnpm.jqh
$ cd ..
$ cd ..
$ cd ..
$ cd svtswb
$ ls
dir fmfnpm
55122 fmfnpm.fdl
dir msfd
30710 rtfwvd.fch
71950 vnwvdmcm.pbz
$ cd fmfnpm
$ ls
318326 rtv.jgn
$ cd ..
$ cd msfd
$ ls
86195 jqnl.hzj
$ cd ..
$ cd ..
$ cd ..
$ cd fwdwq
$ ls
203108 fzb
72579 gcjd.nvp
dir hjfpn
dir qbhn
dir qbnfrhdn
310102 wbdhc
248602 wgdmvslc.ffs
140513 znrlqvmb.thj
dir ztf
$ cd hjfpn
$ ls
dir tfdt
$ cd tfdt
$ ls
202865 chvfjrgz.rdm
$ cd ..
$ cd ..
$ cd qbhn
$ ls
34122 fmfnpm
19676 gfbt.nnc
dir hgdgbrw
dir jlztdcl
93203 jqnl.hzj
dir ppcqqtlj
167401 qzgsswr
$ cd hgdgbrw
$ ls
25297 fzb.vlp
dir nbl
$ cd nbl
$ ls
dir czpfqbp
dir fwdwq
$ cd czpfqbp
$ ls
135628 fzb.tqn
$ cd ..
$ cd fwdwq
$ ls
165198 lrgpwnvv
280597 mcmcn
$ cd ..
$ cd ..
$ cd ..
$ cd jlztdcl
$ ls
188181 bzgzvfjv
116603 fwdvgnqp.fsm
dir fwdwq
277728 fzb
63999 zfrvp.ppd
281623 zgpwjbc.wcj
$ cd fwdwq
$ ls
dir bpt
dir fmrmm
dir fzb
dir qzgsswr
$ cd bpt
$ ls
179418 dhwvhq
313538 njnhmgg
28613 rjc.ncl
267053 sbjbw.shv
$ cd ..
$ cd fmrmm
$ ls
258044 fmfnpm
120035 qcgrnqf.flz
141340 rjc.ncl
dir sps
$ cd sps
$ ls
dir sphjrrvc
$ cd sphjrrvc
$ ls
152609 cwjwbp
$ cd ..
$ cd ..
$ cd ..
$ cd fzb
$ ls
99140 btm
$ cd ..
$ cd qzgsswr
$ ls
25030 cjdltcj.wsw
$ cd ..
$ cd ..
$ cd ..
$ cd ppcqqtlj
$ ls
39010 zcmf.jll
$ cd ..
$ cd ..
$ cd qbnfrhdn
$ ls
282601 cdw.mfn
dir fwdwq
121557 pnqzb.njr
145720 wznfpcn.qhp
$ cd fwdwq
$ ls
152478 fwdvgnqp.fsm
dir fwdwq
dir hhnf
301797 hngcrd.hgt
dir hpj
dir qzgsswr
$ cd fwdwq
$ ls
dir cdb
69663 fmfnpm
97882 fwdvgnqp.fsm
61057 hlg.zhq
115474 rzv.vjv
$ cd cdb
$ ls
193519 djwqhfg.mhl
255813 zptj.cgm
$ cd ..
$ cd ..
$ cd hhnf
$ ls
dir cpzrh
$ cd cpzrh
$ ls
167401 fwdvgnqp.fsm
224559 jqnl.hzj
$ cd ..
$ cd ..
$ cd hpj
$ ls
dir spqljpnt
$ cd spqljpnt
$ ls
268253 fjvqzqt.lmz
63824 lfdrvlc.pww
$ cd ..
$ cd ..
$ cd qzgsswr
$ ls
dir fwdwq
dir hwnmf
dir rbjpbm
dir vhstcbnf
dir zflfcnw
$ cd fwdwq
$ ls
29113 ctr.zjl
130931 jqnl.hzj
111025 zlpmfh
$ cd ..
$ cd hwnmf
$ ls
dir dspbm
260735 fwdwq.qwg
297937 fzb
45087 fzb.drv
19946 jqnl.hzj
51617 nqt.wpr
$ cd dspbm
$ ls
29361 btm
dir fng
127536 lspft.twz
$ cd fng
$ ls
47261 rjc.ncl
$ cd ..
$ cd ..
$ cd ..
$ cd rbjpbm
$ ls
135665 cgfsvqf
dir dbppcf
259879 fwdvgnqp.fsm
143467 qrssbjm.wbh
dir rgvvmf
dir trtgj
$ cd dbppcf
$ ls
72276 jqnl.hzj
$ cd ..
$ cd rgvvmf
$ ls
dir zhldtfj
$ cd zhldtfj
$ ls
248448 hflnwgj
$ cd ..
$ cd ..
$ cd trtgj
$ ls
94714 flfhzc
137893 qqdrglf.rbp
306051 rggqwl
$ cd ..
$ cd ..
$ cd vhstcbnf
$ ls
142465 btm
dir fcpwqj
dir fzb
310858 jmfsz
dir mjjcsbcn
223040 qfnwth.lwq
dir qjfwzdn
dir qzgsswr
156837 rjc.ncl
244544 tvsfmpc.nhz
$ cd fcpwqj
$ ls
75022 thrdj.vnw
$ cd ..
$ cd fzb
$ ls
206638 rjc.ncl
20614 vnq.sfp
277785 zpfjd.vgr
$ cd ..
$ cd mjjcsbcn
$ ls
22117 cscpb.tpw
$ cd ..
$ cd qjfwzdn
$ ls
124990 qzgsswr.pdw
$ cd ..
$ cd qzgsswr
$ ls
dir chptwpbh
dir fwdwq
305061 htbcc.dbl
22511 jqnl.hzj
80135 svvdgvsw
dir vqtnpbn
$ cd chptwpbh
$ ls
55085 fdhzfd.vbl
324100 mcbr.jtn
148699 rtlmz.lvp
$ cd ..
$ cd fwdwq
$ ls
321251 qbqjp.cvs
$ cd ..
$ cd vqtnpbn
$ ls
211984 btm
181182 fwdwq.nsg
212415 lnqbtt.fvs
$ cd ..
$ cd ..
$ cd ..
$ cd zflfcnw
$ ls
183836 jqnl.hzj
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ztf
$ ls
dir vrj
$ cd vrj
$ ls
dir fmfnpm
dir jqmcczl
$ cd fmfnpm
$ ls
52873 sbjbw.shv
$ cd ..
$ cd jqmcczl
$ ls
153695 sbjbw.shv
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd hzpsts
$ ls
dir dsqsfqv
51518 fmfnpm.bhn
dir fszpc
dir fwdwq
dir vrhcvmn
$ cd dsqsfqv
$ ls
221801 hqjwts
$ cd ..
$ cd fszpc
$ ls
dir lrtsfgg
dir nndsvglm
$ cd lrtsfgg
$ ls
316171 lhl.jwm
dir qhph
193371 zbttt.rmd
$ cd qhph
$ ls
71851 zqswjnp.tnc
$ cd ..
$ cd ..
$ cd nndsvglm
$ ls
5209 ddjgsmsm.rhv
dir qzgsswr
$ cd qzgsswr
$ ls
313591 flmfc.pbn
$ cd ..
$ cd ..
$ cd ..
$ cd fwdwq
$ ls
134971 fmfnpm.rcs
$ cd ..
$ cd vrhcvmn
$ ls
95237 fzb.zqw
147520 qzgsswr
dir zrnsdljf
$ cd zrnsdljf
$ ls
120643 sbjbw.shv
$ cd ..
$ cd ..
$ cd ..
$ cd hzrq
$ ls
177154 fwdvgnqp.fsm
127561 fwdwq.hcd
dir vqt
$ cd vqt
$ ls
148546 fwdvgnqp.fsm
108391 rjc.ncl
$ cd ..
$ cd ..
$ cd jzwpjtf
$ ls
46687 fwdwq
$ cd ..
$ cd lmmpmghg
$ ls
dir fwdwq
dir snqlc
dir trtpbbs
dir zlpmfh
$ cd fwdwq
$ ls
dir dzfmw
dir fmfnpm
dir qgdjz
dir qzgsswr
102754 qzgsswr.wns
$ cd dzfmw
$ ls
77275 fmfnpm.qgg
dir fwdwq
$ cd fwdwq
$ ls
132030 fwdvgnqp.fsm
311894 jgvg.lgb
dir mbqw
309968 pwv
185449 qzgsswr.nrz
283317 sbjbw.shv
dir zlpmfh
$ cd mbqw
$ ls
292344 fwdwq.sfb
43979 sbjbw.shv
dir stmb
$ cd stmb
$ ls
164954 fwdwq.gtn
$ cd ..
$ cd ..
$ cd zlpmfh
$ ls
248119 jqnl.hzj
$ cd ..
$ cd ..
$ cd ..
$ cd fmfnpm
$ ls
306459 wrscv
$ cd ..
$ cd qgdjz
$ ls
dir llb
dir pstmbbv
dir qth
dir qzgsswr
dir rqbmzrh
dir ssqj
$ cd llb
$ ls
dir fmfnpm
8728 rjc.ncl
95783 sbjbw.shv
$ cd fmfnpm
$ ls
266355 ddfzzh
$ cd ..
$ cd ..
$ cd pstmbbv
$ ls
303042 fnj.vsr
$ cd ..
$ cd qth
$ ls
155353 chqshnhb.sdm
$ cd ..
$ cd qzgsswr
$ ls
155632 zlpmfh.dhp
$ cd ..
$ cd rqbmzrh
$ ls
202692 rjc.ncl
$ cd ..
$ cd ssqj
$ ls
dir gffp
$ cd gffp
$ ls
34212 rjc.ncl
$ cd ..
$ cd ..
$ cd ..
$ cd qzgsswr
$ ls
324875 btm
275450 fmfnpm.gdr
190317 hwzgndb
251199 mgdsh.tfd
$ cd ..
$ cd ..
$ cd snqlc
$ ls
dir ztsc
$ cd ztsc
$ ls
169368 jqnl.hzj
$ cd ..
$ cd ..
$ cd trtpbbs
$ ls
216211 hhf.hww
101252 htzmmgd
$ cd ..
$ cd zlpmfh
$ ls
312349 qtgpq.jjf
$ cd ..
$ cd ..
$ cd mnw
$ ls
dir fwdwq
dir vbplht
dir zgrwvc
dir zqlqrjgr
$ cd fwdwq
$ ls
dir cgtmbm
dir dpnhgsf
dir lddgc
dir qzgsswr
29505 vffzmhj.jqc
$ cd cgtmbm
$ ls
dir dcw
132840 fwdvgnqp.fsm
dir fwdwq
143226 hwsqv
160598 sbjbw.shv
$ cd dcw
$ ls
dir mdqq
dir nrg
$ cd mdqq
$ ls
230975 fwdvgnqp.fsm
dir fzb
207419 mdwmrb.nqd
dir pfzsg
dir wvtjplcj
$ cd fzb
$ ls
116263 fmfnpm
$ cd ..
$ cd pfzsg
$ ls
211092 chp.lwg
149800 fbffjpz
124128 nqvdq.hhl
$ cd ..
$ cd wvtjplcj
$ ls
82449 lcchd.gsl
$ cd ..
$ cd ..
$ cd nrg
$ ls
32393 fwdvgnqp.fsm
$ cd ..
$ cd ..
$ cd fwdwq
$ ls
63350 zlpmfh.jql
$ cd ..
$ cd ..
$ cd dpnhgsf
$ ls
89773 mchpsf.cwl
$ cd ..
$ cd lddgc
$ ls
56273 cwdjlw
261919 fzb.qbr
49151 vtjc.qqq
$ cd ..
$ cd qzgsswr
$ ls
288467 jqnl.hzj
283162 tfsnpt.cpr
$ cd ..
$ cd ..
$ cd vbplht
$ ls
115696 mfwsbzzm.stf
149951 mnwj.jdn
222362 vtpzr.nnj
dir zlpmfh
$ cd zlpmfh
$ ls
dir dqszzbs
293533 fmfnpm.msw
184810 jsfqrptb.dcl
224988 whmgv
$ cd dqszzbs
$ ls
140899 btm
$ cd ..
$ cd ..
$ cd ..
$ cd zgrwvc
$ ls
dir prmrv
182098 qsljdzsj
16089 rjc.ncl
24028 tpq.dhd
$ cd prmrv
$ ls
261147 rjc.ncl
285206 zlpmfh.hcz
$ cd ..
$ cd ..
$ cd zqlqrjgr
$ ls
dir cjmhdr
16259 sbjbw.shv
$ cd cjmhdr
$ ls
43212 fmfnpm
36273 jqnl.hzj
66482 jtm.hmj
$ cd ..
$ cd ..
$ cd ..
$ cd qzgsswr
$ ls
154596 frglht.mvr
dir fwdwq
dir fzb
150820 rjc.ncl
$ cd fwdwq
$ ls
dir gmthth
121228 mfsmjhr
dir njm
dir tslpnb
$ cd gmthth
$ ls
315665 fzb.gwc
291878 qzgsswr.gss
105707 rjc.ncl
$ cd ..
$ cd njm
$ ls
49199 fwdvgnqp.fsm
17669 jqnl.hzj
$ cd ..
$ cd tslpnb
$ ls
268158 fwdwq
$ cd ..
$ cd ..
$ cd fzb
$ ls
270392 jqnl.hzj
38464 rwfjhg.rqt
19087 wcqmpcrf.tzp
5009 zlpmfh.mmq
$ cd ..
$ cd ..
$ cd zlpmfh
$ ls
dir dfgz
dir fmfnpm
58493 sbjbw.shv
dir tvzqcg
12981 zlpmfh.ggc
$ cd dfgz
$ ls
dir ncrggngr
202192 qdnfl.pvt
$ cd ncrggngr
$ ls
dir fwdwq
$ cd fwdwq
$ ls
246049 btm
123758 vfw
$ cd ..
$ cd ..
$ cd ..
$ cd fmfnpm
$ ls
90482 fwdwq
$ cd ..
$ cd tvzqcg
$ ls
37360 btm
$ cd ..
$ cd ..
$ cd ..
$ cd hchp
$ ls
dir jlcsc
191408 jqnl.hzj
dir lvws
dir zlpmfh
dir zvb
$ cd jlcsc
$ ls
dir pqdwdtm
$ cd pqdwdtm
$ ls
203078 cwtdqrn
169712 fwdvgnqp.fsm
106227 jqnl.hzj
66550 zvmf
$ cd ..
$ cd ..
$ cd lvws
$ ls
210229 tcc.ngf
$ cd ..
$ cd zlpmfh
$ ls
dir vdtv
$ cd vdtv
$ ls
dir fdrv
$ cd fdrv
$ ls
40808 fmfnpm.ssj
$ cd ..
$ cd ..
$ cd ..
$ cd zvb
$ ls
dir cnsddwlg
dir csdz
$ cd cnsddwlg
$ ls
232 jqnl.hzj
dir tcjhhh
$ cd tcjhhh
$ ls
225195 flm.hmb
227170 vvdqw.hwn
$ cd ..
$ cd ..
$ cd csdz
$ ls
dir qzgsswr
$ cd qzgsswr
$ ls
250314 fzb.jhc
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd nsphznf
$ ls
277038 vns.cpc
162685 wvzlgf.dss
dir zldbdl
$ cd zldbdl
$ ls
263687 jqnl.hzj
137415 qzgsswr
$ cd ..
$ cd ..
$ cd phschqg
$ ls
291670 fwdvgnqp.fsm
163036 rjc.ncl
dir zlpmfh
$ cd zlpmfh
$ ls
dir qzgsswr
37363 szdv.chv
$ cd qzgsswr
$ ls
22625 tzhpl
$ cd ..
$ cd ..
$ cd ..
$ cd spfwthmd
$ ls
294499 ltfjljcv
$ cd ..
$ cd wchdqb
$ ls
162305 jtf
215603 zlpmfh
$ cd ..
$ cd zlpmfh
$ ls
dir czb
dir hdm
dir jhqqbqb
dir zmj
$ cd czb
$ ls
17922 fmcljr.qvz
263068 fmlr.sgr
$ cd ..
$ cd hdm
$ ls
dir frbhzrg
dir fwdwq
270049 fwdwq.ldt
297608 glbb
252269 gsdccc.hqw
169953 nlhj.hnc
224702 rjc.ncl
$ cd frbhzrg
$ ls
195491 sbjbw.shv
219888 sgfhs
$ cd ..
$ cd fwdwq
$ ls
dir ggc
dir mvqtp
dir rzvtdvvz
dir vjppb
$ cd ggc
$ ls
150629 fwdwq
144726 fwdwq.csr
18186 fwdwq.jwt
$ cd ..
$ cd mvqtp
$ ls
dir dhp
$ cd dhp
$ ls
132445 btm
$ cd ..
$ cd ..
$ cd rzvtdvvz
$ ls
122759 fwdwq.ptd
$ cd ..
$ cd vjppb
$ ls
286152 jrmjsz.rrq
dir qzgsswr
292661 tnh
$ cd qzgsswr
$ ls
205580 pswmmn.psl
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd jhqqbqb
$ ls
dir mbqc
149539 npgls
156208 nzfphn
dir nzzlwdd
223923 sbjbw.shv
$ cd mbqc
$ ls
11260 lcgnzwcn.pvp
$ cd ..
$ cd nzzlwdd
$ ls
29180 qzgsswr.tcq
232156 zlpmfh.jls
$ cd ..
$ cd ..
$ cd zmj
$ ls
160097 fgt.tfr`
