// Code generated by qtc from "web.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmalert/web.qtpl:1
package main

//line app/vmalert/web.qtpl:3
import (
	"sort"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/tpl"
)

//line app/vmalert/web.qtpl:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmalert/web.qtpl:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmalert/web.qtpl:11
func StreamWelcome(qw422016 *qt422016.Writer) {
//line app/vmalert/web.qtpl:11
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:12
	tpl.StreamHeader(qw422016, "vmalert", navItems)
//line app/vmalert/web.qtpl:12
	qw422016.N().S(`
    <p>
        API:<br>
        `)
//line app/vmalert/web.qtpl:15
	for _, p := range apiLinks {
//line app/vmalert/web.qtpl:15
		qw422016.N().S(`
            `)
//line app/vmalert/web.qtpl:17
		p, doc := p[0], p[1]

//line app/vmalert/web.qtpl:18
		qw422016.N().S(`
        	<a href="`)
//line app/vmalert/web.qtpl:19
		qw422016.E().S(p)
//line app/vmalert/web.qtpl:19
		qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:19
		qw422016.E().S(p)
//line app/vmalert/web.qtpl:19
		qw422016.N().S(`</a> - `)
//line app/vmalert/web.qtpl:19
		qw422016.E().S(doc)
//line app/vmalert/web.qtpl:19
		qw422016.N().S(`<br/>
        `)
//line app/vmalert/web.qtpl:20
	}
//line app/vmalert/web.qtpl:20
	qw422016.N().S(`
    </p>
    `)
//line app/vmalert/web.qtpl:22
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:22
	qw422016.N().S(`
`)
//line app/vmalert/web.qtpl:23
}

//line app/vmalert/web.qtpl:23
func WriteWelcome(qq422016 qtio422016.Writer) {
//line app/vmalert/web.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:23
	StreamWelcome(qw422016)
//line app/vmalert/web.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:23
}

//line app/vmalert/web.qtpl:23
func Welcome() string {
//line app/vmalert/web.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:23
	WriteWelcome(qb422016)
//line app/vmalert/web.qtpl:23
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:23
	return qs422016
//line app/vmalert/web.qtpl:23
}

//line app/vmalert/web.qtpl:25
func StreamListGroups(qw422016 *qt422016.Writer, groups []APIGroup) {
//line app/vmalert/web.qtpl:25
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:26
	tpl.StreamHeader(qw422016, "Groups", navItems)
//line app/vmalert/web.qtpl:26
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:27
	if len(groups) > 0 {
//line app/vmalert/web.qtpl:27
		qw422016.N().S(`
        `)
//line app/vmalert/web.qtpl:29
		rOk := make(map[string]int)
		rNotOk := make(map[string]int)
		for _, g := range groups {
			for _, r := range g.AlertingRules {
				if r.LastError != "" {
					rNotOk[g.Name]++
				} else {
					rOk[g.Name]++
				}
			}
			for _, r := range g.RecordingRules {
				if r.LastError != "" {
					rNotOk[g.Name]++
				} else {
					rOk[g.Name]++
				}
			}
		}

//line app/vmalert/web.qtpl:47
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
        `)
//line app/vmalert/web.qtpl:50
		for _, g := range groups {
//line app/vmalert/web.qtpl:50
			qw422016.N().S(`
              <div class="group-heading`)
//line app/vmalert/web.qtpl:51
			if rNotOk[g.Name] > 0 {
//line app/vmalert/web.qtpl:51
				qw422016.N().S(` alert-danger`)
//line app/vmalert/web.qtpl:51
			}
//line app/vmalert/web.qtpl:51
			qw422016.N().S(`"  data-bs-target="rules-`)
//line app/vmalert/web.qtpl:51
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:51
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line app/vmalert/web.qtpl:52
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:52
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line app/vmalert/web.qtpl:53
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:53
			qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:53
			qw422016.E().S(g.Name)
//line app/vmalert/web.qtpl:53
			if g.Type != "prometheus" {
//line app/vmalert/web.qtpl:53
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:53
				qw422016.E().S(g.Type)
//line app/vmalert/web.qtpl:53
				qw422016.N().S(`)`)
//line app/vmalert/web.qtpl:53
			}
//line app/vmalert/web.qtpl:53
			qw422016.N().S(` (every `)
//line app/vmalert/web.qtpl:53
			qw422016.E().S(g.Interval)
//line app/vmalert/web.qtpl:53
			qw422016.N().S(`)</a>
                 `)
//line app/vmalert/web.qtpl:54
			if rNotOk[g.Name] > 0 {
//line app/vmalert/web.qtpl:54
				qw422016.N().S(`<span class="badge bg-danger" title="Number of rules withs status Error">`)
//line app/vmalert/web.qtpl:54
				qw422016.N().D(rNotOk[g.Name])
//line app/vmalert/web.qtpl:54
				qw422016.N().S(`</span> `)
//line app/vmalert/web.qtpl:54
			}
//line app/vmalert/web.qtpl:54
			qw422016.N().S(`
                <span class="badge bg-success" title="Number of rules withs status Ok">`)
//line app/vmalert/web.qtpl:55
			qw422016.N().D(rOk[g.Name])
//line app/vmalert/web.qtpl:55
			qw422016.N().S(`</span>
                <p class="fs-6 fw-lighter">`)
//line app/vmalert/web.qtpl:56
			qw422016.E().S(g.File)
//line app/vmalert/web.qtpl:56
			qw422016.N().S(`</p>
                `)
//line app/vmalert/web.qtpl:57
			if len(g.ExtraFilterLabels) > 0 {
//line app/vmalert/web.qtpl:57
				qw422016.N().S(`
                    <div class="fs-6 fw-lighter">Extra filter labels
                    `)
//line app/vmalert/web.qtpl:59
				for k, v := range g.ExtraFilterLabels {
//line app/vmalert/web.qtpl:59
					qw422016.N().S(`
                            <span class="float-left badge bg-primary">`)
//line app/vmalert/web.qtpl:60
					qw422016.E().S(k)
//line app/vmalert/web.qtpl:60
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:60
					qw422016.E().S(v)
//line app/vmalert/web.qtpl:60
					qw422016.N().S(`</span>
                    `)
//line app/vmalert/web.qtpl:61
				}
//line app/vmalert/web.qtpl:61
				qw422016.N().S(`
                    </div>
                `)
//line app/vmalert/web.qtpl:63
			}
//line app/vmalert/web.qtpl:63
			qw422016.N().S(`
            </div>
            <div class="collapse" id="rules-`)
//line app/vmalert/web.qtpl:65
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:65
			qw422016.N().S(`">
                <table class="table table-striped table-hover table-sm">
                    <thead>
                        <tr>
                            <th scope="col">Rule</th>
                            <th scope="col" title="Shows if rule's execution ended with error">Error</th>
                            <th scope="col" title="How many samples were produced by the rule">Samples</th>
                            <th scope="col" title="How many seconds ago rule was executed">Updated</th>
                        </tr>
                    </thead>
                    <tbody>
                    `)
//line app/vmalert/web.qtpl:76
			for _, ar := range g.AlertingRules {
//line app/vmalert/web.qtpl:76
				qw422016.N().S(`
                        <tr`)
//line app/vmalert/web.qtpl:77
				if ar.LastError != "" {
//line app/vmalert/web.qtpl:77
					qw422016.N().S(` class="alert-danger"`)
//line app/vmalert/web.qtpl:77
				}
//line app/vmalert/web.qtpl:77
				qw422016.N().S(`>
                            <td>
                                <b>alert:</b> `)
//line app/vmalert/web.qtpl:79
				qw422016.E().S(ar.Name)
//line app/vmalert/web.qtpl:79
				qw422016.N().S(` (for: `)
//line app/vmalert/web.qtpl:79
				qw422016.E().V(ar.For)
//line app/vmalert/web.qtpl:79
				qw422016.N().S(`)<br>
                                <code><pre>`)
//line app/vmalert/web.qtpl:80
				qw422016.E().S(ar.Expression)
//line app/vmalert/web.qtpl:80
				qw422016.N().S(`</pre></code><br>
                                `)
//line app/vmalert/web.qtpl:81
				if len(ar.Labels) > 0 {
//line app/vmalert/web.qtpl:81
					qw422016.N().S(` <b>Labels:</b>`)
//line app/vmalert/web.qtpl:81
				}
//line app/vmalert/web.qtpl:81
				qw422016.N().S(`
                                `)
//line app/vmalert/web.qtpl:82
				for k, v := range ar.Labels {
//line app/vmalert/web.qtpl:82
					qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:83
					qw422016.E().S(k)
//line app/vmalert/web.qtpl:83
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:83
					qw422016.E().S(v)
//line app/vmalert/web.qtpl:83
					qw422016.N().S(`</span>
                                `)
//line app/vmalert/web.qtpl:84
				}
//line app/vmalert/web.qtpl:84
				qw422016.N().S(`
                            </td>
                            <td><div class="error-cell">`)
//line app/vmalert/web.qtpl:86
				qw422016.E().S(ar.LastError)
//line app/vmalert/web.qtpl:86
				qw422016.N().S(`</div></td>
                            <td>`)
//line app/vmalert/web.qtpl:87
				qw422016.N().D(ar.LastSamples)
//line app/vmalert/web.qtpl:87
				qw422016.N().S(`</td>
                            <td>`)
//line app/vmalert/web.qtpl:88
				qw422016.N().FPrec(time.Since(ar.LastExec).Seconds(), 3)
//line app/vmalert/web.qtpl:88
				qw422016.N().S(`s ago</td>
                        </tr>
                    `)
//line app/vmalert/web.qtpl:90
			}
//line app/vmalert/web.qtpl:90
			qw422016.N().S(`
                    `)
//line app/vmalert/web.qtpl:91
			for _, rr := range g.RecordingRules {
//line app/vmalert/web.qtpl:91
				qw422016.N().S(`
                        <tr>
                            <td>
                                <b>record:</b> `)
//line app/vmalert/web.qtpl:94
				qw422016.E().S(rr.Name)
//line app/vmalert/web.qtpl:94
				qw422016.N().S(`<br>
                                <code><pre>`)
//line app/vmalert/web.qtpl:95
				qw422016.E().S(rr.Expression)
//line app/vmalert/web.qtpl:95
				qw422016.N().S(`</pre></code>
                                `)
//line app/vmalert/web.qtpl:96
				if len(rr.Labels) > 0 {
//line app/vmalert/web.qtpl:96
					qw422016.N().S(` <b>Labels:</b>`)
//line app/vmalert/web.qtpl:96
				}
//line app/vmalert/web.qtpl:96
				qw422016.N().S(`
                                `)
//line app/vmalert/web.qtpl:97
				for k, v := range rr.Labels {
//line app/vmalert/web.qtpl:97
					qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:98
					qw422016.E().S(k)
//line app/vmalert/web.qtpl:98
					qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:98
					qw422016.E().S(v)
//line app/vmalert/web.qtpl:98
					qw422016.N().S(`</span>
                                `)
//line app/vmalert/web.qtpl:99
				}
//line app/vmalert/web.qtpl:99
				qw422016.N().S(`
                            </td>
                            <td><div class="error-cell">`)
//line app/vmalert/web.qtpl:101
				qw422016.E().S(rr.LastError)
//line app/vmalert/web.qtpl:101
				qw422016.N().S(`</div></td>
                            <td>`)
//line app/vmalert/web.qtpl:102
				qw422016.N().D(rr.LastSamples)
//line app/vmalert/web.qtpl:102
				qw422016.N().S(`</td>
                            <td>`)
//line app/vmalert/web.qtpl:103
				qw422016.N().FPrec(time.Since(rr.LastExec).Seconds(), 3)
//line app/vmalert/web.qtpl:103
				qw422016.N().S(`s ago</td>
                        </tr>
                    `)
//line app/vmalert/web.qtpl:105
			}
//line app/vmalert/web.qtpl:105
			qw422016.N().S(`
                 </tbody>
                </table>
            </div>
        `)
//line app/vmalert/web.qtpl:109
		}
//line app/vmalert/web.qtpl:109
		qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:111
	} else {
//line app/vmalert/web.qtpl:111
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line app/vmalert/web.qtpl:115
	}
//line app/vmalert/web.qtpl:115
	qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:117
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:117
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:119
}

//line app/vmalert/web.qtpl:119
func WriteListGroups(qq422016 qtio422016.Writer, groups []APIGroup) {
//line app/vmalert/web.qtpl:119
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:119
	StreamListGroups(qw422016, groups)
//line app/vmalert/web.qtpl:119
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:119
}

//line app/vmalert/web.qtpl:119
func ListGroups(groups []APIGroup) string {
//line app/vmalert/web.qtpl:119
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:119
	WriteListGroups(qb422016, groups)
//line app/vmalert/web.qtpl:119
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:119
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:119
	return qs422016
//line app/vmalert/web.qtpl:119
}

//line app/vmalert/web.qtpl:122
func StreamListAlerts(qw422016 *qt422016.Writer, groupAlerts []GroupAlerts) {
//line app/vmalert/web.qtpl:122
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:123
	tpl.StreamHeader(qw422016, "Alerts", navItems)
//line app/vmalert/web.qtpl:123
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:124
	if len(groupAlerts) > 0 {
//line app/vmalert/web.qtpl:124
		qw422016.N().S(`
         <a class="btn btn-primary" role="button" onclick="collapseAll()">Collapse All</a>
         <a class="btn btn-primary" role="button" onclick="expandAll()">Expand All</a>
         `)
//line app/vmalert/web.qtpl:127
		for _, ga := range groupAlerts {
//line app/vmalert/web.qtpl:127
			qw422016.N().S(`
            `)
//line app/vmalert/web.qtpl:128
			g := ga.Group

//line app/vmalert/web.qtpl:128
			qw422016.N().S(`
            <div class="group-heading alert-danger" data-bs-target="rules-`)
//line app/vmalert/web.qtpl:129
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:129
			qw422016.N().S(`">
                <span class="anchor" id="group-`)
//line app/vmalert/web.qtpl:130
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:130
			qw422016.N().S(`"></span>
                <a href="#group-`)
//line app/vmalert/web.qtpl:131
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:131
			qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:131
			qw422016.E().S(g.Name)
//line app/vmalert/web.qtpl:131
			if g.Type != "prometheus" {
//line app/vmalert/web.qtpl:131
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:131
				qw422016.E().S(g.Type)
//line app/vmalert/web.qtpl:131
				qw422016.N().S(`)`)
//line app/vmalert/web.qtpl:131
			}
//line app/vmalert/web.qtpl:131
			qw422016.N().S(`</a>
                <span class="badge bg-danger" title="Number of active alerts">`)
//line app/vmalert/web.qtpl:132
			qw422016.N().D(len(ga.Alerts))
//line app/vmalert/web.qtpl:132
			qw422016.N().S(`</span>
                <br>
                <p class="fs-6 fw-lighter">`)
//line app/vmalert/web.qtpl:134
			qw422016.E().S(g.File)
//line app/vmalert/web.qtpl:134
			qw422016.N().S(`</p>
            </div>
            `)
//line app/vmalert/web.qtpl:137
			var keys []string
			alertsByRule := make(map[string][]*APIAlert)
			for _, alert := range ga.Alerts {
				if len(alertsByRule[alert.RuleID]) < 1 {
					keys = append(keys, alert.RuleID)
				}
				alertsByRule[alert.RuleID] = append(alertsByRule[alert.RuleID], alert)
			}
			sort.Strings(keys)

//line app/vmalert/web.qtpl:146
			qw422016.N().S(`
            <div class="collapse" id="rules-`)
//line app/vmalert/web.qtpl:147
			qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:147
			qw422016.N().S(`">
                `)
//line app/vmalert/web.qtpl:148
			for _, ruleID := range keys {
//line app/vmalert/web.qtpl:148
				qw422016.N().S(`
                    `)
//line app/vmalert/web.qtpl:150
				defaultAR := alertsByRule[ruleID][0]
				var labelKeys []string
				for k := range defaultAR.Labels {
					labelKeys = append(labelKeys, k)
				}
				sort.Strings(labelKeys)

//line app/vmalert/web.qtpl:156
				qw422016.N().S(`
                    <br>
                    <b>alert:</b> `)
//line app/vmalert/web.qtpl:158
				qw422016.E().S(defaultAR.Name)
//line app/vmalert/web.qtpl:158
				qw422016.N().S(` (`)
//line app/vmalert/web.qtpl:158
				qw422016.N().D(len(alertsByRule[ruleID]))
//line app/vmalert/web.qtpl:158
				qw422016.N().S(`)
                     | <span><a target="_blank" href="`)
//line app/vmalert/web.qtpl:159
				qw422016.E().S(defaultAR.SourceLink)
//line app/vmalert/web.qtpl:159
				qw422016.N().S(`">Source</a></span>
                    <br>
                    <b>expr:</b><code><pre>`)
//line app/vmalert/web.qtpl:161
				qw422016.E().S(defaultAR.Expression)
//line app/vmalert/web.qtpl:161
				qw422016.N().S(`</pre></code>
                    <table class="table table-striped table-hover table-sm">
                        <thead>
                            <tr>
                                <th scope="col">Labels</th>
                                <th scope="col">State</th>
                                <th scope="col">Active at</th>
                                <th scope="col">Value</th>
                                <th scope="col">Link</th>
                            </tr>
                        </thead>
                        <tbody>
                        `)
//line app/vmalert/web.qtpl:173
				for _, ar := range alertsByRule[ruleID] {
//line app/vmalert/web.qtpl:173
					qw422016.N().S(`
                            <tr>
                                <td>
                                    `)
//line app/vmalert/web.qtpl:176
					for _, k := range labelKeys {
//line app/vmalert/web.qtpl:176
						qw422016.N().S(`
                                        <span class="ms-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:177
						qw422016.E().S(k)
//line app/vmalert/web.qtpl:177
						qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:177
						qw422016.E().S(ar.Labels[k])
//line app/vmalert/web.qtpl:177
						qw422016.N().S(`</span>
                                    `)
//line app/vmalert/web.qtpl:178
					}
//line app/vmalert/web.qtpl:178
					qw422016.N().S(`
                                </td>
                                <td><span class="badge `)
//line app/vmalert/web.qtpl:180
					if ar.State == "firing" {
//line app/vmalert/web.qtpl:180
						qw422016.N().S(`bg-danger`)
//line app/vmalert/web.qtpl:180
					} else {
//line app/vmalert/web.qtpl:180
						qw422016.N().S(` bg-warning text-dark`)
//line app/vmalert/web.qtpl:180
					}
//line app/vmalert/web.qtpl:180
					qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:180
					qw422016.E().S(ar.State)
//line app/vmalert/web.qtpl:180
					qw422016.N().S(`</span></td>
                                <td>`)
//line app/vmalert/web.qtpl:181
					qw422016.E().S(ar.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line app/vmalert/web.qtpl:181
					qw422016.N().S(`</td>
                                <td>`)
//line app/vmalert/web.qtpl:182
					qw422016.E().S(ar.Value)
//line app/vmalert/web.qtpl:182
					qw422016.N().S(`</td>
                                <td>
                                    <a href="/`)
//line app/vmalert/web.qtpl:184
					qw422016.E().S(g.ID)
//line app/vmalert/web.qtpl:184
					qw422016.N().S(`/`)
//line app/vmalert/web.qtpl:184
					qw422016.E().S(ar.ID)
//line app/vmalert/web.qtpl:184
					qw422016.N().S(`/status">Details</a>
                                </td>
                            </tr>
                        `)
//line app/vmalert/web.qtpl:187
				}
//line app/vmalert/web.qtpl:187
				qw422016.N().S(`
                     </tbody>
                    </table>
                `)
//line app/vmalert/web.qtpl:190
			}
//line app/vmalert/web.qtpl:190
			qw422016.N().S(`
            </div>
            <br>
        `)
//line app/vmalert/web.qtpl:193
		}
//line app/vmalert/web.qtpl:193
		qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:195
	} else {
//line app/vmalert/web.qtpl:195
		qw422016.N().S(`
        <div>
            <p>No items...</p>
        </div>
    `)
//line app/vmalert/web.qtpl:199
	}
//line app/vmalert/web.qtpl:199
	qw422016.N().S(`

    `)
//line app/vmalert/web.qtpl:201
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:201
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:203
}

//line app/vmalert/web.qtpl:203
func WriteListAlerts(qq422016 qtio422016.Writer, groupAlerts []GroupAlerts) {
//line app/vmalert/web.qtpl:203
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:203
	StreamListAlerts(qw422016, groupAlerts)
//line app/vmalert/web.qtpl:203
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:203
}

//line app/vmalert/web.qtpl:203
func ListAlerts(groupAlerts []GroupAlerts) string {
//line app/vmalert/web.qtpl:203
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:203
	WriteListAlerts(qb422016, groupAlerts)
//line app/vmalert/web.qtpl:203
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:203
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:203
	return qs422016
//line app/vmalert/web.qtpl:203
}

//line app/vmalert/web.qtpl:205
func StreamAlert(qw422016 *qt422016.Writer, alert *APIAlert) {
//line app/vmalert/web.qtpl:205
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:206
	tpl.StreamHeader(qw422016, "", navItems)
//line app/vmalert/web.qtpl:206
	qw422016.N().S(`
    `)
//line app/vmalert/web.qtpl:208
	var labelKeys []string
	for k := range alert.Labels {
		labelKeys = append(labelKeys, k)
	}
	sort.Strings(labelKeys)

	var annotationKeys []string
	for k := range alert.Annotations {
		annotationKeys = append(annotationKeys, k)
	}
	sort.Strings(annotationKeys)

//line app/vmalert/web.qtpl:219
	qw422016.N().S(`
    <div class="display-6 pb-3 mb-3">`)
//line app/vmalert/web.qtpl:220
	qw422016.E().S(alert.Name)
//line app/vmalert/web.qtpl:220
	qw422016.N().S(`<span class="ms-2 badge `)
//line app/vmalert/web.qtpl:220
	if alert.State == "firing" {
//line app/vmalert/web.qtpl:220
		qw422016.N().S(`bg-danger`)
//line app/vmalert/web.qtpl:220
	} else {
//line app/vmalert/web.qtpl:220
		qw422016.N().S(` bg-warning text-dark`)
//line app/vmalert/web.qtpl:220
	}
//line app/vmalert/web.qtpl:220
	qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:220
	qw422016.E().S(alert.State)
//line app/vmalert/web.qtpl:220
	qw422016.N().S(`</span></div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Active at
        </div>
        <div class="col">
          `)
//line app/vmalert/web.qtpl:227
	qw422016.E().S(alert.ActiveAt.Format("2006-01-02T15:04:05Z07:00"))
//line app/vmalert/web.qtpl:227
	qw422016.N().S(`
        </div>
      </div>
      </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Expr
        </div>
        <div class="col">
          <code><pre>`)
//line app/vmalert/web.qtpl:237
	qw422016.E().S(alert.Expression)
//line app/vmalert/web.qtpl:237
	qw422016.N().S(`</pre></code>
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Labels
        </div>
        <div class="col">
           `)
//line app/vmalert/web.qtpl:247
	for _, k := range labelKeys {
//line app/vmalert/web.qtpl:247
		qw422016.N().S(`
                <span class="m-1 badge bg-primary">`)
//line app/vmalert/web.qtpl:248
		qw422016.E().S(k)
//line app/vmalert/web.qtpl:248
		qw422016.N().S(`=`)
//line app/vmalert/web.qtpl:248
		qw422016.E().S(alert.Labels[k])
//line app/vmalert/web.qtpl:248
		qw422016.N().S(`</span>
          `)
//line app/vmalert/web.qtpl:249
	}
//line app/vmalert/web.qtpl:249
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Annotations
        </div>
        <div class="col">
           `)
//line app/vmalert/web.qtpl:259
	for _, k := range annotationKeys {
//line app/vmalert/web.qtpl:259
		qw422016.N().S(`
                <b>`)
//line app/vmalert/web.qtpl:260
		qw422016.E().S(k)
//line app/vmalert/web.qtpl:260
		qw422016.N().S(`:</b><br>
                <p>`)
//line app/vmalert/web.qtpl:261
		qw422016.E().S(alert.Annotations[k])
//line app/vmalert/web.qtpl:261
		qw422016.N().S(`</p>
          `)
//line app/vmalert/web.qtpl:262
	}
//line app/vmalert/web.qtpl:262
	qw422016.N().S(`
        </div>
      </div>
    </div>
    <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Group
        </div>
        <div class="col">
           <a target="_blank" href="/groups#group-`)
//line app/vmalert/web.qtpl:272
	qw422016.E().S(alert.GroupID)
//line app/vmalert/web.qtpl:272
	qw422016.N().S(`">`)
//line app/vmalert/web.qtpl:272
	qw422016.E().S(alert.GroupID)
//line app/vmalert/web.qtpl:272
	qw422016.N().S(`</a>
        </div>
      </div>
    </div>
     <div class="container border-bottom p-2">
      <div class="row">
        <div class="col-2">
          Source link
        </div>
        <div class="col">
           <a target="_blank" href="`)
//line app/vmalert/web.qtpl:282
	qw422016.E().S(alert.SourceLink)
//line app/vmalert/web.qtpl:282
	qw422016.N().S(`">Link</a>
        </div>
      </div>
    </div>
    `)
//line app/vmalert/web.qtpl:286
	tpl.StreamFooter(qw422016)
//line app/vmalert/web.qtpl:286
	qw422016.N().S(`

`)
//line app/vmalert/web.qtpl:288
}

//line app/vmalert/web.qtpl:288
func WriteAlert(qq422016 qtio422016.Writer, alert *APIAlert) {
//line app/vmalert/web.qtpl:288
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/web.qtpl:288
	StreamAlert(qw422016, alert)
//line app/vmalert/web.qtpl:288
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/web.qtpl:288
}

//line app/vmalert/web.qtpl:288
func Alert(alert *APIAlert) string {
//line app/vmalert/web.qtpl:288
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/web.qtpl:288
	WriteAlert(qb422016, alert)
//line app/vmalert/web.qtpl:288
	qs422016 := string(qb422016.B)
//line app/vmalert/web.qtpl:288
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/web.qtpl:288
	return qs422016
//line app/vmalert/web.qtpl:288
}
