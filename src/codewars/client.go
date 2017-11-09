package codewars

import (
	"errors"
	"codewars/model"
)

type MessageType byte

const (
	_ MessageType = iota
	GameOver
	AuthenticationToken
	TeamSize
	ProtocolVersion
	GameContext
	PlayerContext
	Move
)

const Version int = 1

/**
 * Стратегия --- интерфейс, содержащий описание методов искусственного интеллекта армии.
 * Каждая пользовательская стратегия должна реализовывать этот интерфейс.
 * Может отсутствовать в некоторых языковых пакетах, если язык не поддерживает интерфейсы.
 */
type Strategy interface {
	/**
	 * Основной метод стратегии, осуществляющий управление армией. Вызывается каждый тик.
	 */
	Move(*model.Move)

	NewGame(*model.Game)
	/**
	 * @param world Текущее состояние мира.
	 */
	GetWorld() *model.World
}

const DefaultHost = "127.0.0.1"
const DefaultPort = "31001"
const DefaultToken = "0000000000000000"

type CodeWars struct {
	*AiCup
}

func Start(s Strategy, args ...string) {
	var host, port, token string

	if len(args) == 3 {
		host, port, token = args[0], args[1], args[2]
	} else {
		host, port, token = DefaultHost, DefaultPort, DefaultToken
	}

	cli := &CodeWars{new(AiCup)}

	if err := cli.Dial(host, port); err == nil {
		defer cli.Close()
		cli.writeToken(token)
		cli.writeProtoVersion(Version)
		cli.ReadTeamSize()

		s.NewGame(cli.readGame())
		m := new(model.Move)
		for cli.readContext(s.GetWorld()) == nil {
			cli.writeMove(m)
			m.Reset()
		}
	} else {
		panic(err)
	}
}

func (c *CodeWars) readGame() *model.Game {
	c.ensureMessageType(GameContext)

	if c.readBool() {
		return &model.Game{
			RandomSeed:                             c.readInt64(),
			TickCount:                              c.readInt(),
			WorldWidth:                             c.readFloat64(),
			WorldHeight:                            c.readFloat64(),
			FogOfWarEnabled:                        c.readBool(),
			VictoryScore:                           c.readInt(),
			FacilityCaptureScore:                   c.readInt(),
			VehicleEliminationScore:                c.readInt(),
			ActionDetectionInterval:                c.readInt(),
			BaseActionCount:                        c.readInt(),
			AdditionalActionCountPerControlCenter:  c.readInt(),
			MaxUnitGroup:                           c.readInt(),
			TerrainWeatherMapColumnCount:           c.readInt(),
			TerrainWeatherMapRowCount:              c.readInt(),
			PlainTerrainVisionFactor:               c.readFloat64(),
			PlainTerrainStealthFactor:              c.readFloat64(),
			PlainTerrainSpeedFactor:                c.readFloat64(),
			SwampTerrainVisionFactor:               c.readFloat64(),
			SwampTerrainStealthFactor:              c.readFloat64(),
			SwampTerrainSpeedFactor:                c.readFloat64(),
			ForestTerrainVisionFactor:              c.readFloat64(),
			ForestTerrainStealthFactor:             c.readFloat64(),
			ForestTerrainSpeedFactor:               c.readFloat64(),
			ClearWeatherVisionFactor:               c.readFloat64(),
			ClearWeatherStealthFactor:              c.readFloat64(),
			ClearWeatherSpeedFactor:                c.readFloat64(),
			CloudWeatherVisionFactor:               c.readFloat64(),
			CloudWeatherStealthFactor:              c.readFloat64(),
			CloudWeatherSpeedFactor:                c.readFloat64(),
			RainWeatherVisionFactor:                c.readFloat64(),
			RainWeatherStealthFactor:               c.readFloat64(),
			RainWeatherSpeedFactor:                 c.readFloat64(),
			VehicleRadius:                          c.readFloat64(),
			TankDurability:                         c.readInt(),
			TankSpeed:                              c.readFloat64(),
			TankVisionRange:                        c.readFloat64(),
			TankGroundAttackRange:                  c.readFloat64(),
			TankAerialAttackRange:                  c.readFloat64(),
			TankGroundDamage:                       c.readInt(),
			TankAerialDamage:                       c.readInt(),
			TankGroundDefence:                      c.readInt(),
			TankAerialDefence:                      c.readInt(),
			TankAttackCooldownTicks:                c.readInt(),
			TankProductionCost:                     c.readInt(),
			IFVDurability:                          c.readInt(),
			IFVSpeed:                               c.readFloat64(),
			IFVVisionRange:                         c.readFloat64(),
			IFVGroundAttackRange:                   c.readFloat64(),
			IFVAerialAttackRange:                   c.readFloat64(),
			IFVGroundDamage:                        c.readInt(),
			IFVAerialDamage:                        c.readInt(),
			IFVGroundDefence:                       c.readInt(),
			IFVAerialDefence:                       c.readInt(),
			IFVAttackCooldownTicks:                 c.readInt(),
			IFVProductionCost:                      c.readInt(),
			ARRVDurability:                         c.readInt(),
			ARRVSpeed:                              c.readFloat64(),
			ARRVVisionRange:                        c.readFloat64(),
			ARRVGroundDefence:                      c.readInt(),
			ARRVAerialDefence:                      c.readInt(),
			ARRVProductionCost:                     c.readInt(),
			ARRVRepairRange:                        c.readFloat64(),
			ARRVRepairSpeed:                        c.readFloat64(),
			HelicopterDurability:                   c.readInt(),
			HelicopterSpeed:                        c.readFloat64(),
			HelicopterVisionRange:                  c.readFloat64(),
			HelicopterGroundAttackRange:            c.readFloat64(),
			HelicopterAerialAttackRange:            c.readFloat64(),
			HelicopterGroundDamage:                 c.readInt(),
			HelicopterAerialDamage:                 c.readInt(),
			HelicopterGroundDefence:                c.readInt(),
			HelicopterAerialDefence:                c.readInt(),
			HelicopterAttackCooldownTicks:          c.readInt(),
			HelicopterProductionCost:               c.readInt(),
			FighterDurability:                      c.readInt(),
			FighterSpeed:                           c.readFloat64(),
			FighterVisionRange:                     c.readFloat64(),
			FighterGroundAttackRange:               c.readFloat64(),
			FighterAerialAttackRange:               c.readFloat64(),
			FighterGroundDamage:                    c.readInt(),
			FighterAerialDamage:                    c.readInt(),
			FighterGroundDefence:                   c.readInt(),
			FighterAerialDefence:                   c.readInt(),
			FighterAttackCooldownTicks:             c.readInt(),
			FighterProductionCost:                  c.readInt(),
			MaxFacilityCapturePoints:               c.readFloat64(),
			FacilityCapturePointsPerVehiclePerTick: c.readFloat64(),
			FacilityWidth:                          c.readFloat64(),
			FacilityHeight:                         c.readFloat64(),
		}
	}

	return nil
}

var ErrGameOver = errors.New("game over")

func (c *CodeWars) readContext(w *model.World) error {
	switch c.readOpcode() {
	case GameOver:
		return ErrGameOver
	case PlayerContext:
		if c.readBool() {
			c.readPlayer(w.Player)
			c.readWorld(w)
		}
		return nil
	default:
		return ErrWrongType
	}
}

func (c *CodeWars) readPlayer(fn func(id int64) *model.Player) {
	switch c.readByte() {
	case 0:
		return
	case 127:
		c.readInt64() // consume id, no changes
	default:
		p := fn(c.readInt64())
		p.Me = c.readBool()
		p.StrategyCrashed = c.readBool()
		p.Score = c.readInt()
		p.RemainingActionCooldownTicks = c.readInt()
	}
}

func (c *CodeWars) readWorld(w *model.World) {
	if c.readBool() {
		w.TickIndex = c.readInt()
		w.TickCount = c.readInt()
		w.Width = c.readFloat64()
		w.Height = c.readFloat64()
		w.LineSize = int(w.Height) / model.TileSize

		c.readPlayers(w)
		c.readVehicles(w)       // New
		c.readVehiclesUpdate(w) // Updates

		if len(w.Land) == 0 {
			w.Land = make([]model.LandType, int(w.Width*w.Height)/(model.TileSize*model.TileSize))

			c.readTerrains(w)
			c.readWeather(w)
		}

		c.readFacilities(w)
	}
}

func (c *CodeWars) writeMove(m *model.Move) {
	c.writeOpcode(Move)

	c.writeBool(true)

	c.writeByte(byte(m.Action))
	c.writeInt(m.Group)
	c.writeFloat64(m.Left)
	c.writeFloat64(m.Top)
	c.writeFloat64(m.Right)
	c.writeFloat64(m.Bottom)
	c.writeFloat64(m.X)
	c.writeFloat64(m.Y)
	c.writeFloat64(m.Angle)
	c.writeFloat64(m.MaxSpeed)
	c.writeFloat64(m.MaxAngularSpeed)
	c.writeByte(byte(m.VehicleType))
	c.writeInt64(m.FacilityId)

	c.flush()
}

func (c *CodeWars) readWeather(w *model.World) {
	idx := 0
	for i := c.readInt(); i > 0; i-- {
		for j := c.readInt(); j > 0; j-- {
			w.Land[idx] = w.Land[idx] & ^(model.Land_Rain | model.Land_Cloud)
			switch c.readByte() {
			case byte(model.Weather_Rain):
				w.Land[idx] |= model.Land_Rain
			case byte(model.Weather_Cloud):
				w.Land[idx] |= model.Land_Cloud
			}
			idx++
		}
	}
}

func (c *CodeWars) readTerrains(w *model.World) {
	idx := 0
	for i := c.readInt(); i > 0; i-- {
		for j := c.readInt(); j > 0; j-- {
			w.Land[idx] = w.Land[idx] & ^(model.Land_Swamp | model.Land_Forest)
			switch c.readByte() {
			case byte(model.Terrain_Swamp):
				w.Land[idx] |= model.Land_Swamp
			case byte(model.Terrain_Forest):
				w.Land[idx] |= model.Land_Forest
			}
			idx++
		}
	}
}

func (c *CodeWars) readFacility(fn func(id int64) *model.Facility) {
	switch c.readByte() {
	case 0, 127:
		break
	default:
		f := fn(c.readInt64())
		f.FacilityType = model.FacilityType(c.readByte())
		f.OwnerPlayerId = c.readInt64()
		f.Left = c.readFloat64()
		f.Top = c.readFloat64()
		f.CapturePoints = c.readFloat64()
		f.VehicleType = model.VehicleType(c.readByte())
		f.ProductionProgress = c.readInt()
	}
}

func (c *CodeWars) readVehicleUpdate(fn func(id int64) *model.Vehicle) {
	if c.readBool() {
		v := fn(c.readInt64())
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Durability = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()
		v.Update()
	}
}

func (c *CodeWars) readVehicle(fn func(id int64) *model.Vehicle) {
	if c.readBool() {
		v := fn(c.readInt64())
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Radius = c.readFloat64()
		v.PlayerId = c.readInt64()
		v.Durability = c.readInt()
		v.MaxDurability = c.readInt()
		v.MaxSpeed = c.readFloat64()
		v.VisionRange = c.readFloat64()
		v.SquaredVisionRange = c.readFloat64()
		v.GroundAttackRange = c.readFloat64()
		v.SquaredGroundAttackRange = c.readFloat64()
		v.AerialAttackRange = c.readFloat64()
		v.SquaredAerialAttackRange = c.readFloat64()
		v.GroundDamage = c.readInt()
		v.AerialDamage = c.readInt()
		v.GroundDefence = c.readInt()
		v.AerialDefence = c.readInt()
		v.AttackCooldownTicks = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.VehicleType = model.VehicleType(c.readByte())
		v.Aerial = c.readBool()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()
	}
}

func (c *CodeWars) readVehiclesUpdate(w *model.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readVehicleUpdate(w.Vehicle)
	}
}

func (c *CodeWars) readFacilities(w *model.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readFacility(w.Facility)
	}
}

func (c *CodeWars) readVehicles(w *model.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readVehicle(w.Vehicle)
	}
}

func (c *CodeWars) readPlayers(w *model.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readPlayer(w.Player)
	}
}