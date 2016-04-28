package main

import (
	"bytes"

	"net/http/httptest"
	"testing"

	"net/http"

	"io/ioutil"
)

func TestRequestResponseHandleByTag(t *testing.T) {

	Init_logger(ioutil.Discard, ioutil.Discard)

	dummy_resp := `<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	   <S:Body>
	      <setApplicationResponse xmlns="urn://x-artefacts-it-ru/dob/poltava/application/types/1.0" xmlns:ns2="urn://x-artefacts-it-ru/dob/poltava/common-types/1.5">
	         <applicationID>72938855-8345-44fa-81d6-7ab93a718787</applicationID>
	         <applicantID>
	            <ns2:personID>c3c648b8-a426-46c4-af49-e91277f6e209</ns2:personID>
	         </applicantID>
	         <applicationState>
	            <applicationStateTransitionID>8</applicationStateTransitionID>
	            <applicationStateName>Заявка оформлена и принята на рассмотрение</applicationStateName>
	            <applicationStateTransitionTimestamp>2016-02-05T14:26:02.440+04:00</applicationStateTransitionTimestamp>
	         </applicationState>
	         <applicationContentID>80819aab-5f0f-49d1-95c1-2dc3baf02d9b</applicationContentID>
	      </setApplicationResponse>
	   </S:Body>
	</S:Envelope>`
	mapping = map[string]string{}
	mapping["dummy.resp"] = dummy_resp

	data := bytes.NewBufferString(`  <bodyColorCode/>
            <vehiclePassportNumber> </vehiclePassportNumber>
            <registeringDocumentType/>
            <haveform/>
            <dovtypehave/>
            <docnamehave/>
            <serianumberhave> </serianumberhave>
            <agenthave/>
            <serianumberpolis> </serianumberpolis>
            <agentpolis/>
          </vehicle>
          <deptcode>6359</deptcode>
          <fulldate>2015-11-09</fulldate>
        </vehicleRegistrationForm>
`)
	expected_resp := bytes.NewBufferString(dummy_resp)
	tag := "<deptcode>"
	resp_file := "dummy.resp"
	handle := handleByTagWithTemplateLogged(ByContainsTag, tag, resp_file)
	req, err := http.NewRequest("POST", "/test", data)
	if err != nil {
		t.Errorf("Can't generate request!")
	}
	recorder := httptest.NewRecorder()
	handle.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("handleByTagWithTemplateLogged doesn't work!")
	}
	if recorder.Body.String() != expected_resp.String() {
		t.Errorf("Response is wrong")
		t.Error("recorded -->> " + recorder.Body.String())
		t.Error("expected -->> " + expected_resp.String())
	}
}
