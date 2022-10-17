package sbrdata

import "encoding/xml"

// Smses reflects the backup data from SMS Backup and Restore (Pro)
type Smses struct {
	XMLName    xml.Name `xml:"smses"`
	Text       string   `xml:",chardata"`
	Count      string   `xml:"count,attr"`
	BackupSet  string   `xml:"backup_set,attr"`
	BackupDate string   `xml:"backup_date,attr"`
	Type       string   `xml:"type,attr"`
	Sms        []struct {
		Text          string `xml:",chardata"`
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
	} `xml:"sms"`
	Mms []struct {
		Text          string `xml:",chardata"`
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
		Parts         struct {
			Text string `xml:",chardata"`
			Part []struct {
				Text     string `xml:",chardata"`
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
			} `xml:"part"`
		} `xml:"parts"`
		Addrs struct {
			Text string `xml:",chardata"`
			Addr []struct {
				Text    string `xml:",chardata"`
				Address string `xml:"address,attr"`
				Type    string `xml:"type,attr"`
				Charset string `xml:"charset,attr"`
			} `xml:"addr"`
		} `xml:"addrs"`
	} `xml:"mms"`
}
