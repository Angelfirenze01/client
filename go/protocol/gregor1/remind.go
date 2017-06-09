// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/gregor1/remind.avdl

package gregor1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type GetRemindersArg struct {
	MaxReminders int `codec:"maxReminders" json:"maxReminders"`
}

func (o GetRemindersArg) DeepCopy() GetRemindersArg {
	return GetRemindersArg{
		MaxReminders: o.MaxReminders,
	}
}

type DeleteRemindersArg struct {
	ReminderIDs []ReminderID `codec:"reminderIDs" json:"reminderIDs"`
}

func (o DeleteRemindersArg) DeepCopy() DeleteRemindersArg {
	return DeleteRemindersArg{
		ReminderIDs: (func(x []ReminderID) []ReminderID {
			var ret []ReminderID
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ReminderIDs),
	}
}

type RemindInterface interface {
	// getReminders gets the reminders waiting to be sent out as a batch. Get at most
	// maxReminders back.
	GetReminders(context.Context, int) (ReminderSet, error)
	// deleteReminders deletes all of the reminders by ReminderID
	DeleteReminders(context.Context, []ReminderID) error
}

func RemindProtocol(i RemindInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.remind",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getReminders": {
				MakeArg: func() interface{} {
					ret := make([]GetRemindersArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetRemindersArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetRemindersArg)(nil), args)
						return
					}
					ret, err = i.GetReminders(ctx, (*typedArgs)[0].MaxReminders)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"deleteReminders": {
				MakeArg: func() interface{} {
					ret := make([]DeleteRemindersArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DeleteRemindersArg)
					if !ok {
						err = rpc.NewTypeError((*[]DeleteRemindersArg)(nil), args)
						return
					}
					err = i.DeleteReminders(ctx, (*typedArgs)[0].ReminderIDs)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemindClient struct {
	Cli rpc.GenericClient
}

// getReminders gets the reminders waiting to be sent out as a batch. Get at most
// maxReminders back.
func (c RemindClient) GetReminders(ctx context.Context, maxReminders int) (res ReminderSet, err error) {
	__arg := GetRemindersArg{MaxReminders: maxReminders}
	err = c.Cli.Call(ctx, "gregor.1.remind.getReminders", []interface{}{__arg}, &res)
	return
}

// deleteReminders deletes all of the reminders by ReminderID
func (c RemindClient) DeleteReminders(ctx context.Context, reminderIDs []ReminderID) (err error) {
	__arg := DeleteRemindersArg{ReminderIDs: reminderIDs}
	err = c.Cli.Call(ctx, "gregor.1.remind.deleteReminders", []interface{}{__arg}, nil)
	return
}
