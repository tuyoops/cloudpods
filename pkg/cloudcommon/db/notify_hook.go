// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"

	"yunion.io/x/onecloud/pkg/mcclient"
)

var (
	updateNotifyHook updateNotifyHookFunc
	createNotifyHook createNotifyHookFunc
	deleteNotifyHook deleteNotifyHookFunc
)

type updateNotifyHookFunc func(ctx context.Context, userCred mcclient.TokenCredential, obj IModel)
type createNotifyHookFunc func(ctx context.Context, userCred mcclient.TokenCredential, obj IModel)
type deleteNotifyHookFunc func(ctx context.Context, userCred mcclient.TokenCredential, obj IModel)

func SetUpdateNotifyHook(f updateNotifyHookFunc) {
	if updateNotifyHook != nil {
		panic("updateNotifyHook already set")
	}
	updateNotifyHook = f
}

func SetCreateNotifyHook(f createNotifyHookFunc) {
	if createNotifyHook != nil {
		panic("createNotifyHook already set")
	}
	createNotifyHook = f
}

func SetDeleteNotifyHook(f deleteNotifyHookFunc) {
	if deleteNotifyHook != nil {
		panic("deleteNotifyHook already set")
	}
	deleteNotifyHook = f
}

func CallUpdateNotifyHook(ctx context.Context, userCred mcclient.TokenCredential, obj IModel) {
	if updateNotifyHook == nil {
		return
	}
	updateNotifyHook(ctx, userCred, obj)
}

func CallCreateNotifyHook(ctx context.Context, userCred mcclient.TokenCredential, obj IModel) {
	if createNotifyHook == nil {
		return
	}
	createNotifyHook(ctx, userCred, obj)
}

func CallDeleteNotifyHook(ctx context.Context, userCred mcclient.TokenCredential, obj IModel) {
	if deleteNotifyHook == nil {
		return
	}
	deleteNotifyHook(ctx, userCred, obj)
}
