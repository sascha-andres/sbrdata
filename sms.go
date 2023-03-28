package sbrdata

import "encoding/xml"

type (
	PartData interface {
		GetSeq() string
		GetCt() string
		GetName() string
		GetChset() string
		GetCd() string
		GetFn() string
		GetCid() string
		GetCl() string
		GetCttS() string
		GetCttT() string
		GetAttrText() string
	}

	// Part is on part
	Part struct {
		Seq      string `xml:"seq,attr"`
		Ct       string `xml:"ct,attr"`
		Name     string `xml:"name,attr"`
		Chset    string `xml:"chset,attr"`
		Cd       string `xml:"cd,attr"`
		Fn       string `xml:"fn,attr"`
		Cid      string `xml:"cid,attr"`
		Cl       string `xml:"cl,attr"`
		CttS     string `xml:"ctt_s,attr"`
		CttT     string `xml:"ctt_t,attr"`
		AttrText string `xml:"text,attr"`
	}

	PartsData interface {
		GetPart() []Part
	}

	// Parts lists the parts of which the MMS is composed
	Parts struct {
		Part []Part `xml:"part"`
	}

	SMSData interface {
		GetProtocol() string
		GetAddress() string
		GetDate() string
		GetType() string
		GetSubject() string
		GetBody() string
		GetToa() string
		GetScToa() string
		GetServiceCenter() string
		GetRead() string
		GetStatus() string
		GetLocked() string
		GetDateSent() string
		GetSubID() string
		GetReadableDate() string
		GetContactName() string
	}

	// SMS represents a simple short message
	SMS struct {
		Protocol      string `xml:"protocol,attr"`
		Address       string `xml:"address,attr"`
		Date          string `xml:"date,attr"`
		Type          string `xml:"type,attr"`
		Subject       string `xml:"subject,attr"`
		Body          string `xml:"body,attr"`
		Toa           string `xml:"toa,attr"`
		ScToa         string `xml:"sc_toa,attr"`
		ServiceCenter string `xml:"service_center,attr"`
		Read          string `xml:"read,attr"`
		Status        string `xml:"status,attr"`
		Locked        string `xml:"locked,attr"`
		DateSent      string `xml:"date_sent,attr"`
		SubID         string `xml:"sub_id,attr"`
		ReadableDate  string `xml:"readable_date,attr"`
		ContactName   string `xml:"contact_name,attr"`
	}

	AddrData interface {
		GetText() string
		GetAddress() string
		GetType() string
		GetCharset() string
	}

	// Addr is one address
	Addr struct {
		Text    string `xml:",chardata"`
		Address string `xml:"address,attr"`
		Type    string `xml:"type,attr"`
		Charset string `xml:"charset,attr"`
	}

	AddrsData interface {
		GetText() string
		GetAddr() []Addr
	}

	// Addrs is an address list
	Addrs struct {
		Text string `xml:",chardata"`
		Addr []Addr `xml:"addr"`
	}

	MMSData interface {
		GetDate() string
		GetSnippet() string
		GetBlockType() string
		GetCtT() string
		GetSource() string
		GetMsgBox() string
		GetAddress() string
		GetSubCs() string
		GetPreviewType() string
		GetMxID() string
		GetRetrSt() string
		GetDTm() string
		GetExp() string
		GetLocked() string
		GetMID() string
		GetOutTime() string
		GetRetrTxt() string
		GetDateSent() string
		GetRead() string
		GetRptA() string
		GetCtCls() string
		GetTimed() string
		GetPri() string
		GetSubID() string
		GetSyncState() string
		GetRespTxt() string
		GetCtL() string
		GetSimID() string
		GetDRpt() string
		GetMarker() string
		GetFileID() string
		GetID() string
		GetPreviewDataTs() string
		GetMType() string
		GetMxExtension() string
		GetRr() string
		GetFavoriteDate() string
		GetSub() string
		GetReadStatus() string
		GetDateMsPart() string
		GetSeen() string
		GetBindID() string
		GetMxIDV2() string
		GetAdvancedSeen() string
		GetRespSt() string
		GetTextOnly() string
		GetNeedDownload() string
		GetSt() string
		GetRetrTxtCs() string
		GetMSize() string
		GetMxStatus() string
		GetTrID() string
		GetMxType() string
		GetDeleted() string
		GetMCls() string
		GetV() string
		GetAccount() string
		GetPreviewData() string
		GetReadableDate() string
		GetContactName() string
		GetParts() Parts
		GetAddrs() Addrs
	}

	// MMS is a multi media message
	MMS struct {
		Date          string `xml:"date,attr"`
		Snippet       string `xml:"snippet,attr"`
		BlockType     string `xml:"block_type,attr"`
		CtT           string `xml:"ct_t,attr"`
		Source        string `xml:"source,attr"`
		MsgBox        string `xml:"msg_box,attr"`
		Address       string `xml:"address,attr"`
		SubCs         string `xml:"sub_cs,attr"`
		PreviewType   string `xml:"preview_type,attr"`
		MxID          string `xml:"mx_id,attr"`
		RetrSt        string `xml:"retr_st,attr"`
		DTm           string `xml:"d_tm,attr"`
		Exp           string `xml:"exp,attr"`
		Locked        string `xml:"locked,attr"`
		MID           string `xml:"m_id,attr"`
		OutTime       string `xml:"out_time,attr"`
		RetrTxt       string `xml:"retr_txt,attr"`
		DateSent      string `xml:"date_sent,attr"`
		Read          string `xml:"read,attr"`
		RptA          string `xml:"rpt_a,attr"`
		CtCls         string `xml:"ct_cls,attr"`
		Timed         string `xml:"timed,attr"`
		Pri           string `xml:"pri,attr"`
		SubID         string `xml:"sub_id,attr"`
		SyncState     string `xml:"sync_state,attr"`
		RespTxt       string `xml:"resp_txt,attr"`
		CtL           string `xml:"ct_l,attr"`
		SimID         string `xml:"sim_id,attr"`
		DRpt          string `xml:"d_rpt,attr"`
		Marker        string `xml:"marker,attr"`
		FileID        string `xml:"file_id,attr"`
		ID            string `xml:"_id,attr"`
		PreviewDataTs string `xml:"preview_data_ts,attr"`
		MType         string `xml:"m_type,attr"`
		MxExtension   string `xml:"mx_extension,attr"`
		Rr            string `xml:"rr,attr"`
		FavoriteDate  string `xml:"favorite_date,attr"`
		Sub           string `xml:"sub,attr"`
		ReadStatus    string `xml:"read_status,attr"`
		DateMsPart    string `xml:"date_ms_part,attr"`
		Seen          string `xml:"seen,attr"`
		BindID        string `xml:"bind_id,attr"`
		MxIDV2        string `xml:"mx_id_v2,attr"`
		AdvancedSeen  string `xml:"advanced_seen,attr"`
		RespSt        string `xml:"resp_st,attr"`
		TextOnly      string `xml:"text_only,attr"`
		NeedDownload  string `xml:"need_download,attr"`
		St            string `xml:"st,attr"`
		RetrTxtCs     string `xml:"retr_txt_cs,attr"`
		MSize         string `xml:"m_size,attr"`
		MxStatus      string `xml:"mx_status,attr"`
		TrID          string `xml:"tr_id,attr"`
		MxType        string `xml:"mx_type,attr"`
		Deleted       string `xml:"deleted,attr"`
		MCls          string `xml:"m_cls,attr"`
		V             string `xml:"v,attr"`
		Account       string `xml:"account,attr"`
		PreviewData   string `xml:"preview_data,attr"`
		ReadableDate  string `xml:"readable_date,attr"`
		ContactName   string `xml:"contact_name,attr"`
		Parts         Parts  `xml:"parts"`
		Addrs         Addrs  `xml:"addrs"`
	}

	MessageData interface {
		GetText() string
		GetCount() string
		GetBackupSet() string
		GetBackupDate() string
		GetType() string
		GetSms() []SMS
		GetMms() []MMS
	}

	// Messages reflects the backup data from SMS Backup and Restore (Pro)
	Messages struct {
		XMLName    xml.Name `xml:"smses"`
		Text       string   `xml:",chardata"`
		Count      string   `xml:"count,attr"`
		BackupSet  string   `xml:"backup_set,attr"`
		BackupDate string   `xml:"backup_date,attr"`
		Type       string   `xml:"type,attr"`
		Sms        []SMS    `xml:"sms"`
		Mms        []MMS    `xml:"mms"`
	}
)

func (p Part) GetSeq() string {
	return p.Seq
}

func (p Part) GetCt() string {
	return p.Ct
}

func (p Part) GetName() string {
	return p.Name
}

func (p Part) GetChset() string {
	return p.Chset
}

func (p Part) GetCd() string {
	return p.Cd
}

func (p Part) GetFn() string {
	return p.Fn
}

func (p Part) GetCid() string {
	return p.Cid
}

func (p Part) GetCl() string {
	return p.Cl
}

func (p Part) GetCttS() string {
	return p.CttS
}

func (p Part) GetCttT() string {
	return p.CttT
}

func (p Part) GetAttrText() string {
	return p.AttrText
}

func (p Parts) GetPart() []Part {
	if p.Part == nil {
		return make([]Part, 0, 0)
	}
	return p.Part
}

func (S SMS) GetProtocol() string {
	return S.Protocol
}

func (S SMS) GetAddress() string {
	return S.Address
}

func (S SMS) GetDate() string {
	return S.Date
}

func (S SMS) GetType() string {
	return S.Type
}

func (S SMS) GetSubject() string {
	return S.Subject
}

func (S SMS) GetBody() string {
	return S.Body
}

func (S SMS) GetToa() string {
	return S.Toa
}

func (S SMS) GetScToa() string {
	return S.ScToa
}

func (S SMS) GetServiceCenter() string {
	return S.ServiceCenter
}

func (S SMS) GetRead() string {
	return S.Read
}

func (S SMS) GetStatus() string {
	return S.Status
}

func (S SMS) GetLocked() string {
	return S.Locked
}

func (S SMS) GetDateSent() string {
	return S.DateSent
}

func (S SMS) GetSubID() string {
	return S.SubID
}

func (S SMS) GetReadableDate() string {
	return S.ReadableDate
}

func (S SMS) GetContactName() string {
	return S.ContactName
}

func (a Addr) GetText() string {
	return a.Text
}

func (a Addr) GetAddress() string {
	return a.Address
}

func (a Addr) GetType() string {
	return a.Type
}

func (a Addr) GetCharset() string {
	return a.Charset
}

func (a Addrs) GetText() string {
	return a.Text
}

func (a Addrs) GetAddr() []Addr {
	if a.Addr == nil {
		return make([]Addr, 0, 0)
	}
	return a.Addr
}

func (M MMS) GetDate() string {
	return M.Date
}

func (M MMS) GetSnippet() string {
	return M.Snippet
}

func (M MMS) GetBlockType() string {
	return M.BlockType
}

func (M MMS) GetCtT() string {
	return M.CtT
}

func (M MMS) GetSource() string {
	return M.Source
}

func (M MMS) GetMsgBox() string {
	return M.MsgBox
}

func (M MMS) GetAddress() string {
	return M.Address
}

func (M MMS) GetSubCs() string {
	return M.SubCs
}

func (M MMS) GetPreviewType() string {
	return M.PreviewType
}

func (M MMS) GetMxID() string {
	return M.MxID
}

func (M MMS) GetRetrSt() string {
	return M.RetrSt
}

func (M MMS) GetDTm() string {
	return M.DTm
}

func (M MMS) GetExp() string {
	return M.Exp
}

func (M MMS) GetLocked() string {
	return M.Locked
}

func (M MMS) GetMID() string {
	return M.MID
}

func (M MMS) GetOutTime() string {
	return M.OutTime
}

func (M MMS) GetRetrTxt() string {
	return M.RetrTxt
}

func (M MMS) GetDateSent() string {
	return M.DateSent
}

func (M MMS) GetRead() string {
	return M.Read
}

func (M MMS) GetRptA() string {
	return M.RptA
}

func (M MMS) GetCtCls() string {
	return M.CtCls
}

func (M MMS) GetTimed() string {
	return M.Timed
}

func (M MMS) GetPri() string {
	return M.GetPri()
}

func (M MMS) GetSubID() string {
	return M.GetSubID()
}

func (M MMS) GetSyncState() string {
	return M.SyncState
}

func (M MMS) GetRespTxt() string {
	return M.RespTxt
}

func (M MMS) GetCtL() string {
	return M.CtL
}

func (M MMS) GetSimID() string {
	return M.SimID
}

func (M MMS) GetDRpt() string {
	return M.DRpt
}

func (M MMS) GetMarker() string {
	return M.Marker
}

func (M MMS) GetFileID() string {
	return M.FileID
}

func (M MMS) GetID() string {
	return M.ID
}

func (M MMS) GetPreviewDataTs() string {
	return M.PreviewDataTs
}

func (M MMS) GetMType() string {
	return M.MType
}

func (M MMS) GetMxExtension() string {
	return M.MxExtension
}

func (M MMS) GetRr() string {
	return M.Rr
}

func (M MMS) GetFavoriteDate() string {
	return M.FavoriteDate
}

func (M MMS) GetSub() string {
	return M.Sub
}

func (M MMS) GetReadStatus() string {
	return M.ReadStatus
}

func (M MMS) GetDateMsPart() string {
	return M.DateMsPart
}

func (M MMS) GetSeen() string {
	return M.Seen
}

func (M MMS) GetBindID() string {
	return M.BindID
}

func (M MMS) GetMxIDV2() string {
	return M.MxIDV2
}

func (M MMS) GetAdvancedSeen() string {
	return M.AdvancedSeen
}

func (M MMS) GetRespSt() string {
	return M.RespSt
}

func (M MMS) GetTextOnly() string {
	return M.TextOnly
}

func (M MMS) GetNeedDownload() string {
	return M.NeedDownload
}

func (M MMS) GetSt() string {
	return M.St
}

func (M MMS) GetRetrTxtCs() string {
	return M.RetrTxtCs
}

func (M MMS) GetMSize() string {
	return M.MSize
}

func (M MMS) GetMxStatus() string {
	return M.MxStatus
}

func (M MMS) GetTrID() string {
	return M.TrID
}

func (M MMS) GetMxType() string {
	return M.MxType
}

func (M MMS) GetDeleted() string {
	return M.Deleted
}

func (M MMS) GetMCls() string {
	return M.MCls
}

func (M MMS) GetV() string {
	return M.V
}

func (M MMS) GetAccount() string {
	return M.Account
}

func (M MMS) GetPreviewData() string {
	return M.PreviewData
}

func (M MMS) GetReadableDate() string {
	return M.ReadableDate
}

func (M MMS) GetContactName() string {
	return M.ContactName
}

func (M MMS) GetParts() Parts {
	return M.Parts
}

func (M MMS) GetAddrs() Addrs {
	return M.Addrs
}

func (m Messages) GetText() string {
	return m.Text
}

func (m Messages) GetCount() string {
	return m.Count
}

func (m Messages) GetBackupSet() string {
	return m.BackupSet
}

func (m Messages) GetBackupDate() string {
	return m.BackupDate
}

func (m Messages) GetType() string {
	return m.Type
}

func (m Messages) GetSms() []SMS {
	if len(m.Sms) == 0 {
		return make([]SMS, 0, 0)
	}
	return m.Sms
}

func (m Messages) GetMms() []MMS {
	if len(m.Mms) == 0 {
		return make([]MMS, 0, 0)
	}
	return m.Mms
}
