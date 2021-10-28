// Automatically generated by the res2go, do not edit.

package ws_client

import (
    _ "github.com/ying32/govcl/pkgs/winappres"
    "github.com/ying32/govcl/vcl"
)

type TClientForm struct {
    *vcl.TForm
    Label1       *vcl.TLabel
    Label2       *vcl.TLabel
    Label3       *vcl.TLabel
    ServerAddr   *vcl.TEdit
    Label4       *vcl.TLabel
    Label5       *vcl.TLabel
    WSStatus     *vcl.TLabel
    WSId         *vcl.TLabel
    LocalPort    *vcl.TEdit
    WSNote       *vcl.TEdit
    ListView1    *vcl.TListView
    WSClientPort *vcl.TEdit
    Label6       *vcl.TLabel
    RunButton    *vcl.TButton

    //::private::
	TClientFormFields
}

var ClientForm *TClientForm

// Loaded in bytes.
// vcl.Application.CreateForm(&ClientForm)

func NewClientForm(owner vcl.IComponent) (root *TClientForm)  {
    vcl.CreateResForm(owner, &root)
    return
}

var ClientFormBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x0A\x43\x6C\x69\x65\x6E\x74\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xB0\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xC2\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0A\x43\x6C\x69\x65\x6E\x74\x46\x6F\x72\x6D\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xB0\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xC2\x01\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x27\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x15\x05\x57\x69\x64\x74\x68\x02\x6C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x1B\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\xE5\x9C\xB0\xE5\x9D\x80\xE5\x92\x8C\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x55\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x5B\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE9\x93\xBE\xE6\x8E\xA5\xE7\x8A\xB6\xE6\x80\x81\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x3E\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xBE\x00\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x9C\xAC\xE5\x9C\xB0\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAB\xAF\xE5\x8F\xA3\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x0A\x53\x65\x72\x76\x65\x72\x41\x64\x64\x72\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x12\x05\x57\x69\x64\x74\x68\x03\x1A\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x57\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\x99\x00\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE9\x93\xBE\xE6\x8E\xA5\xE5\xA4\x87\xE6\xB3\xA8\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x63\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x78\x05\x57\x69\x64\x74\x68\x02\x30\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\xE9\x93\xBE\xE6\x8E\xA5\x49\x64\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x08\x57\x53\x53\x74\x61\x74\x75\x73\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x01\x03\x54\x6F\x70\x02\x5B\x05\x57\x69\x64\x74\x68\x02\x01\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x04\x57\x53\x49\x64\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x01\x03\x54\x6F\x70\x02\x78\x05\x57\x69\x64\x74\x68\x02\x01\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x09\x4C\x6F\x63\x61\x6C\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xBC\x00\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x06\x57\x53\x4E\x6F\x74\x65\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x96\x00\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x09\x4C\x69\x73\x74\x56\x69\x65\x77\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\xCE\x00\x03\x54\x6F\x70\x03\xE2\x00\x05\x57\x69\x64\x74\x68\x03\xC2\x01\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x03\xE5\xBA\x8F\x07\x56\x69\x73\x69\x62\x6C\x65\x08\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x4D\x65\x74\x68\x6F\x64\x05\x57\x69\x64\x74\x68\x02\x3C\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x50\x61\x74\x68\x05\x57\x69\x64\x74\x68\x03\x96\x00\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x44\x61\x74\x61\x05\x57\x69\x64\x74\x68\x02\x50\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x52\x65\x73\x70\x6F\x6E\x73\x65\x05\x57\x69\x64\x74\x68\x02\x64\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x54\x69\x6D\x65\x05\x57\x69\x64\x74\x68\x02\x64\x00\x01\x09\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x55\x73\x65\x54\x69\x6D\x65\x05\x57\x69\x64\x74\x68\x02\x64\x00\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x05\x54\x45\x64\x69\x74\x0C\x57\x53\x43\x6C\x69\x65\x6E\x74\x50\x6F\x72\x74\x04\x4C\x65\x66\x74\x03\x93\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x39\x05\x57\x69\x64\x74\x68\x02\x50\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x36\x04\x4C\x65\x66\x74\x02\x15\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x3C\x05\x57\x69\x64\x74\x68\x02\x75\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x16\x57\x65\x62\x53\x6F\x63\x6B\x65\x74\xE9\x93\xBE\xE6\x8E\xA5\xE7\xAB\xAF\xE5\x8F\xA3\x3A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x52\x75\x6E\x42\x75\x74\x74\x6F\x6E\x04\x4C\x65\x66\x74\x03\x01\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x39\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE5\x90\xAF\xE5\x8A\xA8\xE9\x93\xBE\xE6\x8E\xA5\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(ClientForm, &ClientFormBytes)