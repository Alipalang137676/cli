بسته api

واردات (
	"encoding/json"
	"آزمایش"
)

func TestPullRequest_ChecksStatus(t *testing.T) {
	pr := PullRequest{}
	محموله := `
	{ "commits": { "nodes": [{ "commit": {
		"statusCheckRollup": {
			"زمینه ها": {
				"گره ها": [
					{ "state": "SUCCESS" },
					{ "state": "PENDING" },
					{ "state": "FAILURE" },
					{ "وضعیت": "IN_PROGRESS",
					  "نتیجه گیری": null }،
					{ "وضعیت": "تکمیل شده",
					  "نتیجه گیری": "موفقیت" }،
					{ "وضعیت": "تکمیل شده",
					  "نتیجه گیری": "شکست" }،
					{ "وضعیت": "تکمیل شده",
					  "نتیجه گیری": "ACTION_REQUIRED" }،
					{ "وضعیت": "تکمیل شده",
					  "نتیجه گیری": "STALE" }
				]
			}
		}
	} }] } }
	`
	err := json.Unmarshal([]byte(payload)، &pr)
	معادله (t، خطا، صفر)

	چک ها:= pr.ChecksStatus()
	معادله (t، چک. مجموع، 8)
	معادله (t، چک‌ها. در انتظار، 3)
	معادله (t، بررسی می کند. شکست، 3)
	معادله (t، چک. پاس، 2)
}Root on package api

import (
	"encoding/json"
	"testing"
)

func TestPullRequest_ChecksStatus(t *testing.T) {
	pr := PullRequest{}
	payload := `
	{ "commits": { "nodes": [{ "commit": {
		"statusCheckRollup": {
			"contexts": {
				"nodes": [
					{ "state": "SUCCESS" },
					{ "state": "PENDING" },
					{ "state": "FAILURE" },
					{ "status": "IN_PROGRESS",
					  "conclusion": null },
					{ "status": "COMPLETED",
					  "conclusion": "SUCCESS" },
					{ "status": "COMPLETED",
					  "conclusion": "FAILURE" },
					{ "status": "COMPLETED",
					  "conclusion": "ACTION_REQUIRED" },
					{ "status": "COMPLETED",
					  "conclusion": "STALE" }
				]
			}
		}
	} }] } }
	`
	err := json.Unmarshal([]byte(payload), &pr)
	eq(t, err, nil)

	checks := pr.ChecksStatus()
	eq(t, checks.Total, 8)
	eq(t, checks.Pending, 3)
	eq(t, checks.Failing, 3)
	eq(t, checks.Passing, 2)
}
