package stable

import (
	"fmt"
	"github.com/zegl/kube-score/scorecard"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ScoreMetaStableAvailable checks if the supplied TypeMeta is an unstable object type, that has a stable(r) replacement
func ScoreMetaStableAvailable(meta metav1.TypeMeta) (score scorecard.TestScore) {
	score.Name = "Stable version"
	score.ID = "stable-version"

	withStable := map[string]map[string]string{
		"extensions/v1beta1": {
			"Deployment": "apps/v1",
			"DaemonSet":  "apps/v1",
		},
		"apps/v1beta1": {
			"Deployment":  "apps/v1",
			"StatefulSet": "apps/v1",
		},
		"apps/v1beta2": {
			"Deployment":  "apps/v1",
			"StatefulSet": "apps/v1",
			"DaemonSet":   "apps/v1",
		},
	}

	if inVersion, ok := withStable[meta.APIVersion]; ok {
		if recommendedVersion, ok := inVersion[meta.Kind]; ok {
			score.Grade = scorecard.GradeWarning
			score.AddComment("",
				fmt.Sprintf("The apiVersion and kind %s/%s is deprecated", meta.APIVersion, meta.Kind),
				fmt.Sprintf("It's recommended to use %s instead", recommendedVersion),
			)
			return
		}
	}

	score.Grade = scorecard.GradeAllOK
	return
}
